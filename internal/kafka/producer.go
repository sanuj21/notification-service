package kafka

import (
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(topic string, message string) {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		log.Fatal("KAFKA_BROKER not set in .env file")
	}

	// Create a new Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaBroker},
		Topic:   topic,
	})

	// Send the message
	err := writer.WriteMessages(nil, kafka.Message{
		Value: []byte(message),
	})
	if err != nil {
		log.Fatal("Failed to write message:", err)
	}
	fmt.Println("Message sent:", message)
}
