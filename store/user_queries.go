package store

const usersCreateQuery = `
INSERT INTO observr_user(
	id,
	username,
	email,
	password
) VALUES (
	:id,
	:username,
	:email,
	:password
) RETURNING created_at, updated_at, deleted_at;
`

const usersEmailExistsQuery = `
SELECT true as exists
FROM observr_user
WHERE LOWER(email) = :email
LIMIT 1;
`

const usersUsernameExistsQuery = `
SELECT true as exists
FROM observr_user
WHERE LOWER(username) = :username
LIMIT 1;
`
