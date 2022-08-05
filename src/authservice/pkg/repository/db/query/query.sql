-- name: GetAuth :one
SELECT * FROM auth WHERE username = $1;

-- name: AddAuth :exec
INSERT INTO auth (id, username, password, email) VALUES ($1, $2, $3, $4);