package main

import (
	"fmt"
	"log"
	"notification-service/internal/kafka"
	"os"
	"sync"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	init()

	topic := "notifications"

	var wg sync.WaitGroup
	wg.Add(2)

	// Start producer
	go func() {
		defer wg.Done()
		kafka.ProduceMessage(topic, "Hello, Kafka!")
	}()

	// Start consumer
	go func() {
		defer wg.Done()
		kafka.ConsumeMessages(topic)
	}()

	wg.Wait()
}
