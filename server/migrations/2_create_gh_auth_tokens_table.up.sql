CREATE TABLE github_authentication_tokens (
	id SERIAL PRIMARY KEY,
	github_user_id TEXT NOT NULL,
	encrypted_access_token TEXT NOT NULL
)
