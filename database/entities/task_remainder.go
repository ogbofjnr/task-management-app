package entities

import (
	"time"
)

type TaskRemainder struct {
	Model
	UUID           string `db:"uuid"`
	Task           Task
	TaskID         int       `db:"task_id"`
	TaskStatus     string    `db:"task_status"`
	RemainderTime  time.Time `db:"reminder_time"`
	ReminderStatus string    `db:"reminder_status"`
}
