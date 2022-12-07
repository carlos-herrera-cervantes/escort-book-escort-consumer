package repositories

import (
	"context"
	"time"

	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"
)

//go:generate mockgen -destination=./mocks/iprofile_status_repository.go -package=mocks --build_flags=--mod=mod . IProfileStatusRepository
type IProfileStatusRepository interface {
	GetByProfileId(ctx context.Context, profileId string) (models.ProfileStatus, error)
	Create(ctx context.Context, profileStatus *models.ProfileStatus) error
	UpdateByProfileId(ctx context.Context, profileId, profileStatusCategoryId string) error
}

type ProfileStatusRepository struct {
	Data *db.PostgresClient
}

func (r *ProfileStatusRepository) GetByProfileId(
	ctx context.Context,
	profileId string,
) (profileStatus models.ProfileStatus, err error) {
	query := "SELECT * FROM profile_status WHERE escort_id = $1;"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, profileId)

	if err = row.Scan(
		&profileStatus.Id,
		&profileStatus.EscortId,
		&profileStatus.ProfileStatusCategoryId,
		&profileStatus.CreatedAt,
		&profileStatus.UpdatedAt,
	); err != nil {
		return profileStatus, err
	}

	return profileStatus, nil
}

func (r *ProfileStatusRepository) Create(ctx context.Context, profileStatus *models.ProfileStatus) error {
	query := "INSERT INTO profile_status VALUES ($1, $2, $3, $4, $5);"
	profileStatus.SetDefaultValues()

	_, err := r.Data.EscortProfileDB.ExecContext(
		ctx,
		query,
		profileStatus.Id,
		profileStatus.EscortId,
		profileStatus.ProfileStatusCategoryId,
		time.Now().UTC(),
		time.Now().UTC())

	if err != nil {
		return err
	}

	return nil
}

func (r *ProfileStatusRepository) UpdateByProfileId(
	ctx context.Context,
	profileId,
	profileStatusCategoryId string,
) error {
	query := `UPDATE profile_status
			  SET profile_status_category_id = $1, updated_at = $2
			  WHERE escort_id = $3;`

	_, err := r.Data.EscortProfileDB.ExecContext(ctx, query, profileStatusCategoryId, time.Now().UTC(), profileId)

	if err != nil {
		return err
	}

	return nil
}
