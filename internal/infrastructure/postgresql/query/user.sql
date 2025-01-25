-- name: FindUser :one
SELECT * FROM users WHERE id = $1;

-- name: SaveUser :one
INSERT INTO users (id, email, created_at) VALUES ($1, $2, $3) RETURNING *;