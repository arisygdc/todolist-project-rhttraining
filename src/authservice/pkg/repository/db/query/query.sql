-- name: Login :one
SELECT * FROM auth WHERE username = $1;