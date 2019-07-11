package main

import (
	"account/internal/application/service"
	"account/internal/domain/validate"
	"account/internal/migrate"
	"account/internal/migrate/migrations"
	"account/internal/repository/mysql"
	"account/internal/transport/http_transport"
	"context"
	"database/sql"
	"github.com/Shopify/sarama"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	saramaEvents "github.com/jmwri/go-events/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {
	address := os.Getenv("LISTEN_ADDRESS")
	dsn := os.Getenv("DSN")
	dsn = dsn + "?multiStatements=true&parseTime=true"

	brokers := os.Getenv("KAFKA_BROKERS")
	brokerList := strings.Split(brokers, ",")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	shutdownCh := make(chan struct{})
	errCh := make(chan error)
	go handleShutdownSignal(shutdownCh)
	go func() {
		select {
		case <-shutdownCh:
			break
		case err := <-errCh:
			log.Printf("fatal error: %s", err)
		}
		cancel()
	}()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = migrateDb(db)
	if err != nil {
		panic(err)
	}

	producer, err := createProducer(brokerList)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := producer.Close()
		if err != nil {
			panic(err)
		}
	}()

	publisher := saramaEvents.NewPublisher(producer)

	r := chi.NewRouter()

	accountRepo := mysql.NewAccountMySQL(db)
	accountRules := validate.NewAccountRules(accountRepo)
	accountValidator := validate.NewAccountValidator(accountRules)
	accountService := service.NewAccountService(accountRepo, accountValidator, publisher)
	http_transport.Bootstrap(r, accountService)

	wg.Add(1)
	go http_transport.Start(address, r, wg, ctx, errCh)

	// doneCh will be closed once wg is done
	doneCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-doneCh:
		// we're finished so start the shutdown
		log.Println("all services finished")
	case <-ctx.Done():
		break
		// break out and wait for shutdown
	}

	log.Println("waiting for shutdown")

	select {
	case <-time.After(time.Second * 10):
		log.Println("killed - took too long to shutdown")
	case <-doneCh:
		log.Println("all services shutdown")
	}
}

func handleShutdownSignal(shutdownCh chan struct{}) {
	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	startedShutdown := false
	for {
		<-quitCh
		if startedShutdown {
			os.Exit(0)
		}
		close(shutdownCh)
		startedShutdown = true
	}
}

func migrateDb(db *sql.DB) error {
	migrationArr := []migrate.Migration{
		&migrations.CreateAccountsTable{},
	}
	return migrate.Migrate(db, migrationArr)
}

func createProducer(brokerList []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Producer.Timeout = time.Second
	return sarama.NewSyncProducer(brokerList, config)
}
