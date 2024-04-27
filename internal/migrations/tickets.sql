-- +goose Up
-- +goose StatementBegin

CREATE TABLE tickets(
    id INT PRIMARY KEY,
    user_id INT NOT NULL,
    event_id INT NOT NULL,
    companion BIT NOT NULL
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE tickets

-- +goose StatementEnd