-- name: GetCat :one
SELECT id, name, age, homeless, img_url, created_at
FROM cats
WHERE id = $1;

-- name: GetCats :many
SELECT id, name, img_url
FROM cats
LIMIT CASE WHEN $1::int = 0 THEN 1000000 ELSE $1::int END;


-- name: SaveCat :one
INSERT INTO cats (name, age, homeless, img_url)
VALUES ($1, $2, $3, $4)
RETURNING id, name, age, homeless, img_url, created_at;