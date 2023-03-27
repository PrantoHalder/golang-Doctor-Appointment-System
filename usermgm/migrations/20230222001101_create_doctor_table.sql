-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS doctordetails (
	id BIGSERIAL,
	userid INT NOT NULL,
    doctortypeid INT NOT NULL,
	degree VARCHAR(100) NOT NULL,
	gender VARCHAR(20) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE (userid),
	PRIMARY KEY (id),
	FOREIGN KEY (doctortypeid) REFERENCES doctortype(id),
	FOREIGN KEY (userid) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctor;
-- +goose StatementEnd
