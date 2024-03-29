-- name: CreateUsers :one
INSERT INTO users (ID, created_at, updated_at, name, api_key)
VALUES($1, $2, $3, $4, 
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByAPIKey :one
Select * from users WHERE api_key=$1;