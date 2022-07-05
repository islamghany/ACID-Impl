-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
    $1,$2,$3
) RETURNING *;


-- name: GetAccount :one
SELECT * from accounts where id = $1 LIMIT 1;

-- name: GetAccoutForUpdate :one
SELECT * FROM accounts WHERE id = $1 lIMIT 1 FOR No key UPDATE;

-- name: GetAccounts :many
SELECT * FROM accounts LIMIT $1 OFFSET $2;


-- name: UpdateAccount :one
UPDATE  accounts
SET balance = $2
where id = $1
RETURNING *;


-- name: DeleteAccount :exec
Delete from accounts where id = $1;