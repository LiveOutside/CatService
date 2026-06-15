-- +goose Up
CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    homeless BOOLEAN NOT NULL DEFAULT false,
    img_url VARCHAR(512) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS cats_feedback (
    id SERIAL PRIMARY KEY,
    quality INT NOT NULL,
    cute INT NOT NULL,
    message TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS cats_feedback;
DROP TABLE IF EXISTS cats;

