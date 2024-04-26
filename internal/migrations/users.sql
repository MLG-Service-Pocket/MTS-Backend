-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
    id INT PRIMARY KEY,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    email VARCHAR NOT NULL
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE users;

-- +goose StatementEnd