package main

import (
	"account/internal/application/handler/account"
	"account/internal/application/service"
	"account/internal/migrate"
	"account/internal/migrate/migrations"
	"account/internal/repository/mysql"
	"account/internal/transport/http_transport"
	"database/sql"
	"errors"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	address := os.Getenv("LISTEN_ADDRESS")
	dsn := os.Getenv("DSN")

	wg := &sync.WaitGroup{}
	shutdownCh := make(chan struct{})
	errCh := make(chan error)
	go handleShutdownSignal(errCh)

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

	r := chi.NewRouter()

	accountRepo := mysql.NewAccountMySQL(db)
	accountService := service.NewAccountService(accountRepo)
	getAccountsHandler := account.NewGetAccountsHandler(accountService)
	getAccountHandler := account.NewGetAccountHandler(accountService)

	r.Get("/", getAccountsHandler.HandleHTTP)
	r.Get("/{id}", getAccountHandler.HandleHTTP)

	go http_transport.Start(address, r, wg, shutdownCh, errCh)

	err = <-errCh
	if err != nil {
		log.Printf("fatal err: %s\n", err)
	}

	log.Println("initiating graceful shutdown")
	close(shutdownCh)

	wg.Wait()
	log.Println("shutdown")
}

func handleShutdownSignal(errCh chan error) {
	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	hit := false
	for {
		<-quitCh
		if hit {
			os.Exit(0)
		}
		if !hit {
			errCh <- errors.New("shutdown signal received")
		}
		hit = true
	}
}

func migrateDb(db *sql.DB) error {
	migrationArr := []migrate.Migration{
		&migrations.CreateAccountsTable{},
	}
	return migrate.Migrate(db, migrationArr)
}
