-- name: CreateAccount :one
INSERT INTO accounts(username, password)
VALUES (@username, @password)
RETURNING id;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = @id;

-- name: GetAccountByUsername :one
SELECT *
FROM accounts
WHERE username = @username;
