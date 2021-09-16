package entities

type Task struct {
	Model
	UUID     string `db:"uuid"`
	Title    string `db:"title"`
	UserID   int    `db:"user_id"`
	Assignee User
	Status   string `db:"status"`
}
