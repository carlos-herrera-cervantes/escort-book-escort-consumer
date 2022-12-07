package main

import (
	"escort-book-escort-consumer/consumers"
	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/handlers"
	"escort-book-escort-consumer/repositories"
)

func main() {
	escortHandler := &handlers.EscortHandler{
		ProfileRepository: &repositories.ProfileRepository{
			Data: db.NewPostgresClient(),
		},
		ProfileStatusRepository: &repositories.ProfileStatusRepository{
			Data: db.NewPostgresClient(),
		},
		ProfileStatusCategoryRepository: &repositories.ProfileStatusCategoryRepository{
			Data: db.NewPostgresClient(),
		},
		NationalityRepository: &repositories.NationalityRepository{
			Data: db.NewPostgresClient(),
		},
	}
	escortProfileConsumer := consumers.EscortProfileConsumer{
		EventHandler: escortHandler,
	}

	escortProfileConsumer.StartConsumer()
}
