-- name: CreateUser :one
INSERT INTO users (
  username,
  hashpassword,
  full_name,
  email,
  isactive
) VALUES (
  $1, $2,$3,$4,$5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;