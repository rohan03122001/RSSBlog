-- +goose Up
ALTER TABLE feeds ADD column last_fetched_at TIMESTAMP;

-- +goose Down

ALTER TABLE feeds DROP column last_fetched_at;