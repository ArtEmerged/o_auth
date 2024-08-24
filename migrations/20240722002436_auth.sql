-- +goose Up
CREATE TYPE status_enum AS ENUM ( 'UNKNOWN', 'ACTIVE', 'BLOCKED', 'DELETED');

CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(256) NOT NULL,
    email varchar(256) NOT NULL UNIQUE,
    pass_hash varchar(64) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    status status_enum NOT NULL DEFAULT 'UNKNOWN',
    role smallint NOT NULL 
);

-- +goose Down
DROP TABLE users;

DROP TYPE status_enum;