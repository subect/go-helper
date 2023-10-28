package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func InitializeKafkaWriter(brokers []string, topic string) *kafka.Writer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return w
}

func PublishMessage(w *kafka.Writer, key, value string) error {
	message := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}

	return w.WriteMessages(context.Background(), message)
}