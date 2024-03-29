-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS doctortype (
	id BIGSERIAL,
	doctortype VARCHAR(20) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL,
	PRIMARY KEY (id),
	UNIQUE (doctortype)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctorType;
-- +goose StatementEnd
