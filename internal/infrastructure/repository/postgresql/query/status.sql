-- name: FindStatus :one
SELECT 
  s.id, 
  s.title,
  s.board_id,
  COALESCE(
    array_agg(ss.next_status) FILTER (WHERE ss.status_id IS NOT NULL),
    ARRAY[]::UUID[]
  ) as next_status,
  COALESCE(
    array_agg(ss.previous_status) FILTER (WHERE ss.status_id IS NOT NULL),
    ARRAY[]::UUID[]
  ) as previous_status,
  s.created_at
FROM statuses s
LEFT JOIN statuses_next_statuses ss ON ss.status_id = s.id
LEFT JOIN statuses_previous_statuses ssp ON ssp.status_id = s.id
WHERE s.id = $1;

-- name: SaveStatus :one
INSERT INTO statuses (id, title, board_id, created_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteStatus :exec
DELETE FROM statuses WHERE id = $1;

-- name: UpdateStatus :exec
UPDATE statuses SET title = $2, board_id = $3, created_at = $4 WHERE id = $1;