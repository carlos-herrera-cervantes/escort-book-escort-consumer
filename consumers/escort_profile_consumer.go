package consumers

import (
	"context"

	"escort-book-escort-consumer/config"
	"escort-book-escort-consumer/handlers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/inconshreveable/log15"
)

type EscortProfileConsumer struct {
	EventHandler handlers.IEventHandler
}

var logger = log.New("consumers")

func (c EscortProfileConsumer) StartConsumer() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  config.InitializeKafka().BootstrapServers,
		"group.id":           config.InitializeKafka().GroupId,
		"auto.offset.reset":  "smallest",
		"enable.auto.commit": true,
	})
	topics := []string{
		config.InitializeKafka().Topics.EscortCreated,
		config.InitializeKafka().Topics.UserActiveAccount,
	}
	_ = consumer.SubscribeTopics(topics, nil)
	run := true

	for run {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
			c.EventHandler.HandleEvent(context.Background(), e)
			logger.Info("Processed message")
		case kafka.PartitionEOF:
			logger.Info("Reached: ", e)
		case kafka.Error:
			logger.Error("Error: ", e)
			run = false
		}
	}

	_ = consumer.Close()
}
