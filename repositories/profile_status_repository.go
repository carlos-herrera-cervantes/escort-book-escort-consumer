package repositories

import (
	"context"
	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"
	"time"
)

type IProfileStatusRepository interface {
	Create(ctx context.Context, profileStatus *models.ProfileStatus) error
}

type ProfileStatusRepository struct {
	Data *db.Data
}

func (r *ProfileStatusRepository) Create(ctx context.Context, profileStatus *models.ProfileStatus) error {
	query := "INSERT INTO profile_status VALUES ($1, $2, $3, $4, $5);"
	profileStatus.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		profileStatus.Id,
		profileStatus.EscortId,
		profileStatus.ProfileStatusCategoryId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return nil
	}

	return nil
}
