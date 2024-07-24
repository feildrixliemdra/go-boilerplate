-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id bigserial PRIMARY KEY ,
    name text NOT NULL,
    email text unique  NOT NULL,
    password text NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp ,
    deleted_at timestamp DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
