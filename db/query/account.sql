-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
    $1,$2,$3
) RETURNING *;


-- name: GetAccount :one
SELECT * from accounts where id = $1;


-- name: GetAccounts :many
SELECT * FROM accounts LIMIT $1 OFFSET $2;


-- name: UpdateAccount :exec
UPDATE  accounts
SET balance = $2
where id = $1;


-- name: DeleteAccount :exec
Delete from accounts where id = $1;