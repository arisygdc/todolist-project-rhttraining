-- name: AddUser :exec
INSERT INTO users (id, first_name, last_name, auth_id) VALUES ($1, $2, $3, $4);

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;