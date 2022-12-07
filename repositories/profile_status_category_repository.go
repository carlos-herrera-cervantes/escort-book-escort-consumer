package repositories

import (
	"context"

	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"
)

//go:generate mockgen -destination=./mocks/iprofile_status_category_repository.go -package=mocks --build_flags=--mod=mod . IProfileStatusCategoryRepository
type IProfileStatusCategoryRepository interface {
	GetOneByName(ctx context.Context, name string) (models.ProfileStatusCategory, error)
}

type ProfileStatusCategoryRepository struct {
	Data *db.PostgresClient
}

func (r *ProfileStatusCategoryRepository) GetOneByName(ctx context.Context, name string) (models.ProfileStatusCategory, error) {
	query := "SELECT * FROM profile_status_category WHERE name = $1"
	row := r.Data.EscortProfileDB.QueryRowContext(ctx, query, name)

	var category models.ProfileStatusCategory
	err := row.Scan(&category.Id, &category.Name, &category.Active, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return models.ProfileStatusCategory{}, err
	}

	return category, nil
}
