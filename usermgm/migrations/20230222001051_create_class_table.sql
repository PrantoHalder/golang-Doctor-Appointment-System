-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS doctor_type (
	id int,
	name VARCHAR(20) NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL,
	PRIMARY KEY (id),
	UNIQUE (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctor_type;
-- +goose StatementEnd
