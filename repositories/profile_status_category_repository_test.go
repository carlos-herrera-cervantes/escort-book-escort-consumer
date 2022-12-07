package repositories

import (
	"context"
	"testing"
	"time"

	"escort-book-escort-consumer/db"

	"github.com/stretchr/testify/assert"
)

func TestProfileStatusCategoryRepositoryGetOneByName(t *testing.T) {
	profileStatusCategoryRepository := ProfileStatusCategoryRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error when record does not exists", func(t *testing.T) {
		_, err := profileStatusCategoryRepository.GetOneByName(context.Background(), "dummy")
		assert.Error(t, err)
	})

	t.Run("Should return category", func(t *testing.T) {
		_, _ = profileStatusCategoryRepository.Data.EscortProfileDB.ExecContext(
			context.Background(),
			"INSERT INTO profile_status_category VALUES ($1, $2, $3, $4, $5);",
			"638d8d5dd605d9f9e2b4e495",
			"Dummy",
			true,
			time.Now().UTC(),
			time.Now().UTC(),
		)
		_, err := profileStatusCategoryRepository.GetOneByName(context.Background(), "Dummy")
		assert.NoError(t, err)

		_, _ = profileStatusCategoryRepository.Data.EscortProfileDB.ExecContext(
			context.Background(),
			"DELETE FROM profile_status_category;",
		)
	})
}
