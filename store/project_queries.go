package store

const projectsCreateQuery = `
INSERT INTO observr_project(
	id,
	name,
	url,
	user_id,
	api_key
) VALUES (
	:id,
	:name,
	:url,
	:user_id,
	:api_key
) RETURNING created_at, updated_at, deleted_at;
`
