-- +goose Up
CREATE TABLE users (
    id serial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    pass_hash text NOT NULL,
    created_at timestamptz  NOT NULL,
    updated_at timestamptz,
    deleted_at timestamptz,
    role integer NOT NULL
);

-- +goose Down
DROP TABLE users;