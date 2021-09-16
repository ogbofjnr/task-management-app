CREATE TABLE task_reminders
(
    id              serial PRIMARY KEY,
    uuid            VARCHAR   NOT NULL,
    task_status     VARCHAR   NOT NULL,
    task_id         INTEGER   NOT NULL,
    reminder_status VARCHAR   NOT NULL,
    reminder_time   timestamp NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    deleted_at      TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_task_reminders_task_id FOREIGN KEY (task_id) REFERENCES tasks (id)
);
CREATE INDEX idx_task_reminders_uuid ON task_reminders (uuid);

CREATE INDEX idx_task_reminders_task_id ON task_reminders (task_id);

