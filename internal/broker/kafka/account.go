package kafka

import (
	"account/internal/domain/account"
	"account/internal/mapper"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

const newAccountTopic = "queue.account.new"

func NewAccountKafka(producer sarama.SyncProducer) *AccountKafka {
	return &AccountKafka{
		Producer: producer,
	}
}

type AccountKafka struct {
	Producer sarama.SyncProducer
}

func (r *AccountKafka) New(a *account.Account) error {
	acc := mapper.AccountToEvent(a)
	encoded, err := json.Marshal(acc)
	if err != nil {
		return err
	}
	partition, offset, err := r.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: newAccountTopic,
		Value: sarama.StringEncoder(encoded),
	})

	if err != nil {
		return err
	}
	log.Printf("published event %s with partition %d offset %s", newAccountTopic, partition, offset)
	return nil
}
