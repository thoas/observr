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
