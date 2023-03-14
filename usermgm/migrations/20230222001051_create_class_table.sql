-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS doctor_type (
	id BIGSERIAL,
	doctor_type TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL,
	PRIMARY KEY (id),
	UNIQUE (doctor_type)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctor_type;
-- +goose StatementEnd
