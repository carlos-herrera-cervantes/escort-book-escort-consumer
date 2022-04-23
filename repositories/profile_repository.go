package repositories

import (
	"context"
	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"
	"time"
)

type IProfileRepository interface {
	Create(ctx context.Context, profile *models.Profile) error
}

type ProfileRepository struct {
	Data *db.Data
}

func (r *ProfileRepository) Create(ctx context.Context, profile *models.Profile) error {
	query := "INSERT INTO profile VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	profile.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
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
