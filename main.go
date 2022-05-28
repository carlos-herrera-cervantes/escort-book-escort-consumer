package main

import (
	"context"
	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/handlers"
	"escort-book-escort-consumer/repositories"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVERS"),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
		"auto.offset.reset": "smallest",
		"enable.auto.commit": true,
	})

	_ = consumer.SubscribeTopics(
		[]string{os.Getenv("KAFKA_ESCORT_TOPIC"), os.Getenv("KAFKA_ACTIVE_ACCOUNT_TOPIC")},
		nil,
	)
	run := true
	handler := &handlers.EscortHandler{
		ProfileRepository: &repositories.ProfileRepository{
			Data: db.New(),
		},
		ProfileStatusRepository: &repositories.ProfileStatusRepository{
			Data: db.New(),
		},
		ProfileStatusCategoryRepository: &repositories.ProfileStatusCategoryRepository{
			Data: db.New(),
		},
		NationalityRepository: &repositories.NationalityRepository{
			Data: db.New(),
		},
	}

	for run == true {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
			handler.ProcessMessage(context.Background(), e)
			log.Println("PROCESSED MESSAGE")
		case kafka.PartitionEOF:
			log.Println("Reached: ", e)
		case kafka.Error:
			log.Println("Error: ", e)
			run = false
		default:
		}
	}

	_ = consumer.Close()
}
