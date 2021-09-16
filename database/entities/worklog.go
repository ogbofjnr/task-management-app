package entities

type Worklog struct {
	Model
	UUID       string `db:"uuid"`
	Task       Task
	TaskID     int    `db:"task_id"`
	LoggedTime int    `db:"logged_time"`
	Status     string `db:"status"`
}
