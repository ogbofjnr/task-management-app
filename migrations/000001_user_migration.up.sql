CREATE TABLE users
(
    id         serial PRIMARY KEY,
    uuid       VARCHAR      NOT NULL,
    email      VARCHAR      NOT NULL,
    first_name VARCHAR      NOT NULL,
    last_name  VARCHAR      NOT NULL,
    status     VARCHAR      NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
CREATE INDEX idx_users_uuid ON users (uuid);

CREATE UNIQUE INDEX idx_users_email ON users (email);