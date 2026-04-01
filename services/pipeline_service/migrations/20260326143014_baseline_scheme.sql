-- +goose Up
CREATE TABLE IF NOT EXISTS pipeline(
    id          uuid PRIMARY KEY,
    name        text NOT NULL,
    description text,
    created_at  timestamptz DEFAULT now(),
    updated_at  timestamptz,
    deleted_at  timestamptz
);

-- +goose Down
DROP TABLE IF EXISTS pipeline;
