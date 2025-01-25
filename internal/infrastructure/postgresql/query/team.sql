-- name: FindTeam :one
SELECT 
  t.id, 
  t.title,
  t.description, 
  t.owner,
  COALESCE(
    array_agg(tu.user_id) FILTER (WHERE tu.team_id IS NOT NULL),
    ARRAY[]::UUID[]
  ) as members,
  t.created_at
FROM teams t
LEFT JOIN teams_users tu ON tu.team_id = t.id
WHERE t.id = $1;

-- name: SaveTeam :one
INSERT INTO teams (id, title, description, owner, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM teams WHERE id = $1;

-- name: UpdateTeam :exec
UPDATE teams SET title = $2, description = $3, owner = $4, created_at = $5 WHERE id = $1;