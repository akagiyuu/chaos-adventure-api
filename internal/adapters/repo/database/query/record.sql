-- name: CreateRecord :exec
INSERT INTO records(account_id, time)
VALUES (@account_id, @time);

-- name: GetAllRecord :many
SELECT 
    (SELECT username FROM accounts WHERE id = r.account_id) as username,
    time,
    created_at
FROM records r;
