-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS doctor_schedule (
	id BIGSERIAL,
	doctorid INT NOT NULL,
	startat TIMESTAMP NOT NULL,
	endat TIMESTAMP NOT NULL,
	workdays JSON NOT NULL,
	address TEXT NOT NULL,
	phone varchar(20) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id),
	UNIQUE (phone),
	FOREIGN KEY (doctorid) REFERENCES doctordetails(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctorschedule;
-- +goose StatementEnd
