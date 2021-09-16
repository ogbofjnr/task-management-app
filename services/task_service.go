package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ogbofjnr/maze/constants"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/database/repositories"
	"github.com/ogbofjnr/maze/errors"
	requests "github.com/ogbofjnr/maze/requests/task"
)

type TaskService struct {
	TaskRepository *repositories.TaskRepository
	db             *sqlx.DB
}

func NewTaskService(TaskRepository *repositories.TaskRepository, db *sqlx.DB) *TaskService {
	return &TaskService{
		TaskRepository: TaskRepository,
		db:             db,
	}
}

func (s *TaskService) CreateTask(request *requests.CreateTaskRequest, user *entities.User) (*entities.Task, error) {
	Task := &entities.Task{}
	Task.UUID = uuid.New().String()
	Task.Title = request.Title
	Task.Status = constants.Todo
	Task.UserID = user.ID
	err := s.TaskRepository.Insert(s.db, Task)
	if err != nil {
		return nil, fmt.Errorf("error creating Task: %w", err)
	}
	return Task, nil
}

func (s *TaskService) GetTask(request *requests.GetTaskRequest) (*entities.Task, error) {
	Task, err := s.TaskRepository.GetByUUID(s.db, request.UUID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving Task: %w", err)
	}
	return Task, nil
}

func (s *TaskService) DeleteTask(request *requests.DeleteTaskRequest) error {
	Task, err := s.TaskRepository.GetByUUID(s.db, request.UUID)
	if err != nil {
		return fmt.Errorf("error retrieving Task: %w", err)
	}
	err = s.TaskRepository.Delete(s.db, Task)
	if err != nil {
		return fmt.Errorf("error deleting Task: %w", err)
	}
	return nil
}

func (s *TaskService) SearchTasks(request *requests.SearchTaskRequest) (*[]entities.Task, error) {
	Tasks, err := s.TaskRepository.Search(s.db, request)
	if err != nil {
		return nil, fmt.Errorf("error retrieving task: %w", err)
	}
	return Tasks, nil
}

func (s *TaskService) SetTaskStatus(request *requests.SetTaskStatus) error {
	tx, err := s.db.Beginx()
	defer tx.Rollback()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	tx.MustExec("LOCK TABLE tasks IN EXCLUSIVE MODE")
	task, err := s.GetTask(&requests.GetTaskRequest{UUID: request.TaskUUID})
	if err != nil {
		return fmt.Errorf("error getting task: %w", err)
	}

	userActiveTasks, err := s.TaskRepository.GetUserActiveTasks(tx, task.UserID)
	if err != nil {
		return fmt.Errorf("error getting user active tasks: %w", err)
	}
	if len(*userActiveTasks) != 0 {
		return errors.UserAlreadyHaveActiveTask
	}
	task.Status = request.Status
	err = s.TaskRepository.UpdateTask(tx, task)
	if err != nil {
		return fmt.Errorf("error retrieving task: %w", err)
	}
	tx.Commit()
	return nil
}

func (s *TaskService) AddWorklog(request *requests.AddWorklogRequest) error {
	task, err := s.GetTask(&requests.GetTaskRequest{UUID: request.TaskUUID})
	if err != nil {
		return fmt.Errorf("error getting task: %w", err)
	}
	Worklog := &entities.Worklog{}
	Worklog.UUID = uuid.New().String()
	Worklog.LoggedTime = request.LoggedTime
	Worklog.Status = constants.Active
	Worklog.TaskID = task.ID
	err = s.TaskRepository.AddWorklog(s.db, Worklog)
	if err != nil {
		return fmt.Errorf("error adding worklog : %w", err)
	}
	return nil
}

func (s *TaskService) SetReminder(request *requests.SetReminderRequest) error {
	task, err := s.GetTask(&requests.GetTaskRequest{UUID: request.TaskUUID})
	if err != nil {
		return fmt.Errorf("error getting task: %w", err)
	}
	taskReminder := &entities.TaskRemainder{}
	taskReminder.UUID = uuid.New().String()
	taskReminder.RemainderTime = request.ReminderTime
	taskReminder.ReminderStatus = constants.Active
	taskReminder.TaskID = task.ID
	taskReminder.TaskStatus = request.TaskStatus
	err = s.TaskRepository.SetReminder(s.db, taskReminder)
	if err != nil {
		return fmt.Errorf("error creating task remider: %w", err)
	}
	return nil
}
