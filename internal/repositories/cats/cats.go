package cats

import (
	"cat_service/internal/data/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCat(ctx context.Context, db *pgxpool.Pool, id int) (models.CatResponse, error) {
	var cat models.CatResponse

	query := `
		SELECT id, name, age, homeless, img_url, created_at
		FROM cats
		WHERE id = $1
	`

	err := db.QueryRow(ctx, query, id).Scan(
		&cat.ID,
		&cat.Name,
		&cat.Age,
		&cat.Homeless,
		&cat.ImageUrl,
		&cat.CreatedAt,
	)

	if err != nil {
		return models.CatResponse{}, fmt.Errorf("failed to scan cat")
	}

	return cat, nil
}

func GetCats(ctx context.Context, db *pgxpool.Pool, limit int) ([]models.CatPreview, error) {
	var (
		cats []models.CatPreview
		rows pgx.Rows
		err  error
	)

	if limit == 0 {
		query := `SELECT id, name, img_url FROM cats`
		rows, err = db.Query(ctx, query)
	} else {
		query := `SELECT id, name, img_url FROM cats LIMIT $1`
		rows, err = db.Query(ctx, query, limit)
	}

	if err != nil {
		return []models.CatPreview{}, fmt.Errorf("failed to retrieve cats")
	}

	for rows.Next() {
		var cat models.CatPreview

		err := rows.Scan(
			&cat.ID,
			&cat.Name,
			&cat.ImageUrl,
		)

		if err != nil {
			return []models.CatPreview{}, fmt.Errorf("failed to scan cat")
		}
		cats = append(cats, cat)
	}

	return cats, nil
}

func SaveCat(ctx context.Context, pool *pgxpool.Pool, cat models.CatRequest) (models.CatResponse, error) {
	var response models.CatResponse

	query := `
		INSERT INTO cats (name, age, homeless, img_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, age, homeless, img_url, created_at
	`

	err := pool.QueryRow(ctx, query, cat.Name, cat.Age, cat.Homeless, cat.ImageUrl).Scan(
		&response.ID,
		&response.Name,
		&response.Age,
		&response.Homeless,
		&response.ImageUrl,
		&response.CreatedAt,
	)
	if err != nil {
		return models.CatResponse{}, fmt.Errorf("insert failed: %w", err)
	}

	return response, nil
}
