-- +goose Up
-- +goose StatementBegin

CREATE TABLE events(
    id INT PRIMARY KEY,
    title VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    location VARCHAR NOT NULL,
    min_price INT NOT NULL,
    closest_start_time TIMESTAMP NOT NULL,
    last_start_time TIMESTAMP NULL,
    tags TEXT[]
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE tickets

-- +goose StatementEnd