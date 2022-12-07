package repositories

import (
	"context"
	"testing"
	"time"

	"escort-book-escort-consumer/db"

	"github.com/stretchr/testify/assert"
)

func TestNationalityRepositoryGetOneByName(t *testing.T) {
	nationalityRepository := NationalityRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error when record does not exists", func(t *testing.T) {
		_, err := nationalityRepository.GetOneByName(context.Background(), "dummy")
		assert.Error(t, err)
	})

	t.Run("Should return nationality", func(t *testing.T) {
		_, _ = nationalityRepository.Data.EscortProfileDB.ExecContext(
			context.Background(),
			"INSERT INTO nationality VALUES ($1, $2, $3, $4, $5);",
			"638d7e883fdf92f13bf965c1",
			"Dummy",
			true,
			time.Now().UTC(),
			time.Now().UTC(),
		)

		_, err := nationalityRepository.GetOneByName(context.Background(), "Dummy")
		assert.NoError(t, err)

		_, _ = nationalityRepository.Data.EscortProfileDB.ExecContext(
			context.Background(),
			"DELETE FROM nationality;",
		)
	})
}
