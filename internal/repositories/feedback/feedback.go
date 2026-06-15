package feedback

import (
	"cat_service/internal/data/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveFeedback(ctx context.Context, db *pgxpool.Pool, feedback models.Feedback) error {
	query := "INSERT INTO cats_feedback (quality, cute, message) VALUES ($1, $2, $3)"
	_, err := db.Exec(ctx, query, feedback.Quality, feedback.Cute, feedback.Message)
	return err
}
