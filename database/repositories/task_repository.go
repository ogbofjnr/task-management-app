package repositories

import (
	sql2 "database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/ogbofjnr/maze/constants"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/errors"
	"github.com/ogbofjnr/maze/pkg/db"
	requests "github.com/ogbofjnr/maze/requests/task"
)

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

// Insert inserts task into db
func (r *TaskRepository) Insert(db db.DB, Task *entities.Task) error {
	_, err := db.NamedExec(`
	INSERT INTO tasks (uuid,title,user_id,status,created_at,updated_at)
	VALUES (:uuid, :title, :user_id, :status, now(), now())`, Task)

	if err != nil {
		return fmt.Errorf("error inserting task to db: %w", err)
	}

	return nil
}

// GetByUUID finds task by UUID
func (r *TaskRepository) GetByUUID(db db.DB, uuid string) (*entities.Task, error) {
	Task := &entities.Task{}
	sqStmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := sqStmt.Select("*").From("tasks").Where("status!=? AND uuid=?", constants.Deleted, uuid).ToSql()

	if err != nil {
		return nil, fmt.Errorf("error query sql: %w", err)
	}

	err = db.Get(Task, sql, args...)

	if err == sql2.ErrNoRows {
		return nil, errors.TaskNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("error selecting Task from db: %w", err)
	}
	return Task, nil
}

// Delete soft deletes task
func (r *TaskRepository) Delete(db db.DB, Task *entities.Task) error {
	_, err := db.Exec(`UPDATE tasks SET status=$1, deleted_at=now() WHERE id=$2`, constants.Deleted, Task.ID)

	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}

	return nil
}

// Search finds tasks by given condition
func (r *TaskRepository) Search(db db.DB, request *requests.SearchTaskRequest) (*[]entities.Task, error) {
	Tasks := &[]entities.Task{}
	sqStmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := sqStmt.Select("*").From("tasks")

	if request.Status != nil {
		builder = builder.Where("status=?", request.Status)
	}

	if request.Title != nil {
		builder = builder.Where("title=?", request.Title)
	}

	sql, args, err := builder.ToSql()

	if err != nil {
		return nil, fmt.Errorf("error query sql: %w", err)
	}

	err = db.Select(Tasks, sql, args...)

	if err != nil {
		return nil, fmt.Errorf("error selecting task from db: %w", err)
	}
	return Tasks, nil
}

// UpdateTask updates task status and title
func (r *TaskRepository) UpdateTask(db db.DB, task *entities.Task) error {
	_, err := db.NamedExec(`UPDATE tasks SET title=:title, status=:status WHERE id=:id`, task)

	if err != nil {
		return fmt.Errorf("error updating task status: %w", err)
	}

	return nil
}

// GetUserActiveTasks retrives user's active tasks
func (r *TaskRepository) GetUserActiveTasks(db db.DB, userID int) (*[]entities.Task, error) {
	Tasks := &[]entities.Task{}
	err := db.Select(Tasks,
		"SELECT FROM tasks t WHERE t.id IN(select t.id FROM tasks tt INNER JOIN users u ON tt.user_id=u.id WHERE u.id=$1 AND tt.status=$2)",
		userID, constants.InProgress)

	if err != nil {
		return nil, fmt.Errorf("error selecting user's tasks from db: %w", err)
	}
	return Tasks, nil
}

// AddWorklog adds worklog record for the task
func (r *TaskRepository) AddWorklog(db db.DB, worklog *entities.Worklog) error {
	_, err := db.NamedExec(`
	INSERT INTO worklogs (uuid,logged_time,task_id,status,created_at,updated_at)
	VALUES (:uuid, :logged_time, :task_id, :status, now(), now())`, worklog)

	if err != nil {
		return fmt.Errorf("error inserting worklogs to db: %w", err)
	}

	return nil
}

// SetReminder sets reminder for the task
func (r *TaskRepository) SetReminder(db db.DB, taskRemainder *entities.TaskRemainder) error {
	_, err := db.NamedExec(`
	INSERT INTO task_reminders (uuid,reminder_time,task_id,task_status,reminder_status,created_at,updated_at)
	VALUES (:uuid, :reminder_time, :task_id, :task_status,:status, now(), now())`, taskRemainder)

	if err != nil {
		return fmt.Errorf("error inserting taskRemainder to db: %w", err)
	}

	return nil
}

// GetTasksForNotification finds reminders to be executed
func (r *TaskRepository) GetTasksForNotification(db db.DB) (*[]entities.Task, error) {
	Tasks := &[]entities.Task{}
	err := db.Select(Tasks,
		`SELECT * FROM tasks t 
					WHERE t.id IN(select tr.task_id from task_reminders tr WHERE tr.reminder_time>now() and tr.reminder_status=$1)`, constants.Active)

	if err != nil {
		return Tasks, fmt.Errorf("error selecting tasks from db: %w", err)
	}
	return Tasks, nil
}
