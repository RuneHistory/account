package events

import (
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
	"log"
)

func NewKafkaDispatcher(producer sarama.SyncProducer) *KafkaDispatcher {
	return &KafkaDispatcher{
		Producer: producer,
	}
}

type KafkaDispatcher struct {
	Producer sarama.SyncProducer
}

func (d *KafkaDispatcher) Dispatch(e Event) error {
	event, ok := e.(KafkaEvent)
	if !ok {
		return errors.New("non-kafka event given")
	}
	encoded, err := json.Marshal(event.Body())
	if err != nil {
		return err
	}
	partition, offset, err := d.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: event.Topic(),
		Value: sarama.StringEncoder(encoded),
	})

	if err != nil {
		return err
	}
	log.Printf("published event %s with partition %d offset %s", event.Topic(), partition, offset)
	return nil
}

type KafkaEvent interface {
	Event
	Topic() string
}
