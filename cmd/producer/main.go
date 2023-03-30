package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChannel := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("Hello World", "teste", producer, []byte("18282"), deliveryChannel)
	go DeliveryReport(deliveryChannel)
	defer producer.Flush(2000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "172.17.0.1:9092",
		"delivery.timeout.ms": 1000,
		"acks":                "all",
		"retries":             10,
		"enable.idempotence":  "true",
	}

	producer, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println("Error creating producer =>", err.Error())
	}

	return producer
}

func Publish(message string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
	messageData := &kafka.Message{
		Value: []byte(message),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	err := producer.Produce(messageData, deliveryChannel)

	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliverychan chan kafka.Event) {
	for event := range deliverychan {
		switch e := event.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				log.Println("Error delivering message =>", e.TopicPartition.Error)
			} else {
				fmt.Println("Message Delivered =>", e.TopicPartition)
			}
		}
	}
}
