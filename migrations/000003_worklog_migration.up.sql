CREATE TABLE worklogs
(
    id          serial PRIMARY KEY,
    uuid        VARCHAR   NOT NULL,
    logged_time VARCHAR   NOT NULL,
    status      VARCHAR   NOT NULL,
    task_id     INT       NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    deleted_at  TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_worklog_task_id FOREIGN KEY(task_id) REFERENCES tasks(id)
);
CREATE INDEX idx_worklogs_uuid ON worklogs (uuid);
CREATE INDEX idx_worklogs_task_id ON worklogs (task_id);

