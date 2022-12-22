-- name: CreateSecret :one
INSERT INTO secrets (
  content,
  creator,
  hashpassword,
  isview
) VALUES (
  $1, $2,$3,$4
)
RETURNING *;

-- name: GetSecret :one
SELECT * FROM secrets
WHERE id = $1 LIMIT 1;

-- name: ListSecrets :many
SELECT * FROM secrets
WHERE creator = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteSecret :exec
DELETE FROM secrets
WHERE id = $1;

-- name: UpdateSecret :one
UPDATE secrets
SET isview = $2
WHERE id = $1
RETURNING *;
