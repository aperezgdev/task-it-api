-- name: FindTask :one
SELECT * FROM tasks WHERE id = $1;

-- name: SaveTask :one
INSERT INTO tasks (id, title, description, creator, asigned, status_id, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: UpdateTask :exec
UPDATE tasks SET title = $2, description = $3, creator = $4, asigned = $5, status_id = $6, created_at = $7 WHERE id = $1;