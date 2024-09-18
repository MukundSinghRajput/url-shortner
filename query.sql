-- name: Add :one
INSERT INTO tiny (
    origin, short
) VALUES (
    $1, $2
)
RETURNING origin, short;

-- name: GetByOrigin :one
SELECT origin, short
FROM tiny
WHERE origin = $1
LIMIT 1;

-- name: GetByShort :one
SELECT origin, short
FROM tiny
WHERE short = $1
LIMIT 1;