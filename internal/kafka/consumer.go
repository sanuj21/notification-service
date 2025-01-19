package kafka

import (
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(topic string) {
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaGroupID := os.Getenv("KAFKA_GROUP_ID")
	if kafkaBroker == "" || kafkaGroupID == "" {
		log.Fatal("KAFKA_BROKER or KAFKA_GROUP_ID not set in .env file")
	}

	// Create a new Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		GroupID: kafkaGroupID,
		Topic:   topic,
	})

	// Start consuming messages
	for {
		message, err := reader.ReadMessage(nil)
		if err != nil {
			log.Fatal("Failed to read message:", err)
		}
		fmt.Println("Message received:", string(message.Value))
	}
}
