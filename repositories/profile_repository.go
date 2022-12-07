package repositories

import (
	"context"
	"time"

	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"
)

//go:generate mockgen -destination=./mocks/iprofile_repository.go -package=mocks --build_flags=--mod=mod . IProfileRepository
type IProfileRepository interface {
	Create(ctx context.Context, profile *models.Profile) error
}

type ProfileRepository struct {
	Data *db.PostgresClient
}

func (r *ProfileRepository) Create(ctx context.Context, profile *models.Profile) error {
	query := "INSERT INTO profile VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	profile.SetDefaultValues()

	_, err := r.Data.EscortProfileDB.ExecContext(
		ctx,
		query,
		profile.Id,
		profile.EscortId,
		"empty",
		"empty",
		profile.Email,
		"empty",
		profile.Gender,
		profile.NationalityId,
		"1900-01-01",
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}
