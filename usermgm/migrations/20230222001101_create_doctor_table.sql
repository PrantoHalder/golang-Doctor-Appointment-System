-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS doctor (
	id BIGSERIAL,
	userid INT NOT NULL,
    doctortypeID INT NOT NULL,
	degree VARCHAR(100) NOT NULL,
	gender VARCHAR(20) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL,

	PRIMARY KEY (id),
	FOREIGN KEY (doctortypeID) REFERENCES doctortype(id),
	FOREIGN KEY (userid) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctor;
-- +goose StatementEnd
