package repositories

import (
	"context"
	"testing"
	"time"

	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"

	"github.com/stretchr/testify/assert"
)

func setUpInitialState() {
	repository := ProfileStatusRepository{
		Data: db.NewPostgresClient(),
	}

	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"INSERT INTO nationality VALUES ($1, $2, $3, $4, $5);",
		"96877f7f-a292-4bed-83d2-9cc9a1d5acc3",
		"Dummy Nationality",
		true,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"INSERT INTO profile_status_category VALUES ($1, $2, $3, $4, $5);",
		"eb248737-4c5e-4826-a358-70fae61089fe",
		"Dummy Category 1",
		true,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"INSERT INTO profile_status_category VALUES ($1, $2, $3, $4, $5);",
		"6f7a5b42-bbc2-4e01-9a43-cf95ca3cf349",
		"Dummy Category 2",
		true,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"INSERT INTO profile VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);",
		"f51e87dc-f839-4165-8ab1-102d09f79a1b",
		"638ee5845788924fb913b58c",
		"empty",
		"empty",
		"test.user2@example.com",
		"empty",
		"Male",
		"96877f7f-a292-4bed-83d2-9cc9a1d5acc3",
		"1994-01-01",
		time.Now().UTC(),
		time.Now().UTC(),
	)
}

func setUpFinalState() {
	repository := ProfileStatusRepository{
		Data: db.NewPostgresClient(),
	}

	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"DELETE FROM profile;",
	)
	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"DELETE FROM profile_status;",
	)
	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"DELETE FROM nationality;",
	)
	_, _ = repository.Data.EscortProfileDB.ExecContext(
		context.Background(),
		"DELETE FROM profile_status_category;",
	)
}

func TestProfileStatusRepositoryCRUD(t *testing.T) {
	profileStatusRepository := ProfileStatusRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should apply CRUD operations", func(t *testing.T) {
		setUpInitialState()

		profileStatus := models.ProfileStatus{
			Id:                      "80f83958-6eda-4cb3-85ff-7e80ac23661b",
			EscortId:                "638ee5845788924fb913b58c",
			ProfileStatusCategoryId: "eb248737-4c5e-4826-a358-70fae61089fe",
		}
		err := profileStatusRepository.Create(context.Background(), &profileStatus)
		assert.NoError(t, err)

		_, err = profileStatusRepository.GetByProfileId(context.Background(), "638ee5845788924fb913b58c")
		assert.NoError(t, err)

		err = profileStatusRepository.UpdateByProfileId(
			context.Background(),
			"638ee5845788924fb913b58c",
			"6f7a5b42-bbc2-4e01-9a43-cf95ca3cf349",
		)
		assert.NoError(t, err)

		setUpFinalState()
	})
}

func TestProfileStatusRepositoryGetByProfileId(t *testing.T) {
	profileStatusRepository := ProfileStatusRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error when query fails", func(t *testing.T) {
		_, err := profileStatusRepository.GetByProfileId(context.Background(), "638eea59d6ad521589a87d74")
		assert.Error(t, err)
	})
}

func TestProfileStatusRepositoryCreate(t *testing.T) {
	profileStatusRepository := ProfileStatusRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error when insert fails", func(t *testing.T) {
		profileStatus := models.ProfileStatus{
			Id:                      "2bf17360-9398-4002-af63-e0b8d0ed3374",
			EscortId:                "638eead9b1dc1bd32487e827",
			ProfileStatusCategoryId: "43d135c7-d0a7-4dcd-9522-f4394aa365c9",
		}
		err := profileStatusRepository.Create(context.Background(), &profileStatus)
		assert.Error(t, err)
	})
}

func TestProfileStatusRepositoryUpdateByProfileId(t *testing.T) {
	profileStatusRepository := ProfileStatusRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error when update fails", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := profileStatusRepository.UpdateByProfileId(
			ctxWithCancel,
			"638eeb82851d1f8d6937373a",
			"c1611db5-c38b-4126-ad4b-89b0af62aef0",
		)
		assert.Error(t, err)
	})
}
