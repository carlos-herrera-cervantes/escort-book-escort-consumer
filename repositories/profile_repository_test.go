package repositories

import (
	"context"
	"testing"
	"time"

	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"

	"github.com/stretchr/testify/assert"
)

func TestProfileRepositoryCreate(t *testing.T) {
	profileRepository := ProfileRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error when insert fails", func(t *testing.T) {
		profile := models.Profile{}
		err := profileRepository.Create(context.Background(), &profile)
		assert.Error(t, err)
	})

	t.Run("Should return nil when the insert is successful", func(t *testing.T) {
		_, _ = profileRepository.Data.EscortProfileDB.ExecContext(
			context.Background(),
			"INSERT INTO nationality VALUES ($1, $2, $3, $4, $5);",
			"bc0acec2-fd96-4879-9e86-e209e5c3a50d",
			"Dummy",
			true,
			time.Now().UTC(),
			time.Now().UTC(),
		)
		profile := models.Profile{
			Id:            "2584833a-bc79-4fcd-94b0-7d522595a66a",
			EscortId:      "638d8555a401a930b03f8dae",
			Email:         "test.user@example.com",
			Gender:        "Male",
			NationalityId: "bc0acec2-fd96-4879-9e86-e209e5c3a50d",
		}
		err := profileRepository.Create(context.Background(), &profile)
		assert.NoError(t, err)

		_, _ = profileRepository.Data.EscortProfileDB.ExecContext(
			context.Background(),
			"DELETE FROM profile;",
		)
		_, _ = profileRepository.Data.EscortProfileDB.ExecContext(
			context.Background(),
			"DELETE FROM nationality;",
		)
	})
}
