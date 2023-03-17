-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS appointment (
	id BIGSERIAL,
	userid INT NOT NULL,
	doctordetailsid INT NOT NULL,
	schduleid INT NOT NULL,
	is_appointed BOOLEAN DEFAULT true,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	FOREIGN KEY (userid) REFERENCES users(id),
	FOREIGN KEY (doctordetailsid) REFERENCES doctordetails(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS appointment;
-- +goose StatementEnd
