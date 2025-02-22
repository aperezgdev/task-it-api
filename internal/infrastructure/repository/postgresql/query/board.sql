-- name: FindBoard :one
SELECT * FROM boards WHERE id = $1;

-- name: SaveBoard :one
INSERT INTO boards (id, title, description, owner, team_id, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: DeleteBoard :exec
DELETE FROM boards WHERE id = $1;

-- name: UpdateBoard :exec 
UPDATE boards SET title = $2, description = $3, owner = $4, team_id = $5, created_at = $6 WHERE id = $1;
