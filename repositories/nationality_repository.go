package repositories

import (
	"context"
	"escort-book-escort-consumer/db"
	"escort-book-escort-consumer/models"
)

type INationalityRepository interface {
	GetOneByName(ctx context.Context, name string) (models.Nationality, error)
}

type NationalityRepository struct {
	Data *db.Data
}

func (r *NationalityRepository) GetOneByName(ctx context.Context, name string) (models.Nationality, error) {
	query := "SELECT id, name FROM nationality WHERE name = $1"
	row := r.Data.DB.QueryRowContext(ctx, query, name)

	var nationality models.Nationality
	err := row.Scan(&nationality.Id, &nationality.Name)

	if err != nil {
		return models.Nationality{}, err
	}

	return nationality, nil
}
