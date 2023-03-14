-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
     id BIGSERIAL NOT NULL,
	 first_name VARCHAR(20) NOT NULL,
	 last_name VARCHAR(20) NOT NULL,
	 username TEXT NOT NULL,
	 email TEXT NOT NULL,
	 password TEXT NOT NULL,
	 role VARCHAR(20) NOT NULL DEFAULT 'user',
	 is_active BOOLEAN DEFAULT TRUE,
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	 deleted_at TIMESTAMP DEFAULT NULL,
	 PRIMARY KEY (id),
	 UNIQUE (username),
	 UNIQUE (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users ;
-- +goose StatementEnd
