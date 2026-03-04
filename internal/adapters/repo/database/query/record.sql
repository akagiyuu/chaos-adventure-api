-- name: CreateRecord :exec
INSERT INTO records(account_id, time)
VALUES (@account_id, @time);

-- name: GetAllRecord :many
SELECT *
FROM records;
