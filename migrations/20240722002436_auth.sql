-- +goose Up
CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(256) NOT NULL,
    email varchar(256) NOT NULL UNIQUE,
    pass_hash varchar(64) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    status varchar(10) NOT NULL DEFAULT 'UNKNOWN',
    role integer NOT NULL
);

-- +goose Down
DROP TABLE users;