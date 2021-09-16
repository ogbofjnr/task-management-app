CREATE TABLE tasks
(
    id         serial PRIMARY KEY,
    uuid       VARCHAR   NOT NULL,
    title      VARCHAR   NOT NULL,
    status     VARCHAR   NOT NULL,
    user_id    INT       NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_tasks_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);
CREATE INDEX idx_tasks_uuid ON tasks (uuid);

CREATE INDEX idx_tasks_user_id ON tasks (user_id);
