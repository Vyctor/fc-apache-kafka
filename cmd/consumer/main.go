package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "172.17.0.1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-consumer-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		fmt.Println("Error creating consumer =>", err.Error())
	}

	topics := []string{"teste"}

	consumer.SubscribeTopics(topics, nil)

	for {
		message, err := consumer.ReadMessage(-1)

		if err != nil {
			fmt.Println("Error consuming message =>", err.Error())

		}

		fmt.Println("Message =>", string(message.Value), "Partition =>", message.TopicPartition.Partition, "Offset =>", message.TopicPartition.Offset)
	}
}
