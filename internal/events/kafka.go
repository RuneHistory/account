package events

import (
	"encoding/json"
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
	event := e.(KafkaEvent)
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
	log.Printf("published event %s with partition %d offset %d", event.Topic(), partition, offset)
	return nil
}

type KafkaEvent interface {
	Event
	Topic() string
}
