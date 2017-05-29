package store

const usersCreateQuery = `
INSERT INTO observr_user(
	id,
	username,
	email,
	password,
	api_key
) VALUES (
	:id,
	:username,
	:email,
	:password,
	:api_key
) RETURNING created_at, updated_at, deleted_at;
`

var usersGetByIDQuery = `
SELECT *
FROM observr_user
WHERE id = :id
  AND deleted_at IS NULL
LIMIT 1;
`

var usersGetByAPIKeyQuery = `
SELECT *
FROM observr_user
WHERE api_key = :api_key
  AND deleted_at IS NULL
LIMIT 1;
`

const usersEmailExistsQuery = `
SELECT true as exists
FROM observr_user
WHERE email = :email
LIMIT 1;
`

const usersUsernameExistsQuery = `
SELECT true as exists
FROM observr_user
WHERE LOWER(username) = :username
LIMIT 1;
`
