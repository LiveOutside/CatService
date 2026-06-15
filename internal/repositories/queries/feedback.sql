-- name: SaveFeedback :exec
INSERT INTO cats_feedback (quality, cute, message)
VALUES ($1, $2, $3);