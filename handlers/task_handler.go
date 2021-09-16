package handlers

import (
	"errors"
	appErrors "github.com/ogbofjnr/maze/errors"
	validator2 "github.com/ogbofjnr/maze/pkg/validator"
	requests "github.com/ogbofjnr/maze/requests/task"
	userRequests "github.com/ogbofjnr/maze/requests/user"
	"github.com/ogbofjnr/maze/responses"
	taskResponses "github.com/ogbofjnr/maze/responses/task"
	"github.com/ogbofjnr/maze/services"
	"go.uber.org/zap"
	"net/http"
)

type TaskHandler struct {
	validator   validator2.Validator
	TaskService *services.TaskService
	UserService *services.UserService
	logger      *zap.Logger
}

func NewTaskHandler(
	validator validator2.Validator,
	TaskService *services.TaskService,
	userService *services.UserService,
	logger *zap.Logger,
) *TaskHandler {
	return &TaskHandler{
		validator:   validator,
		TaskService: TaskService,
		UserService: userService,
		logger:      logger,
	}
}

// Create creates new task
func (u *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateTaskRequest(r)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	err = u.validator.Validate(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors(u.validator.GetErrors(err)), u.logger)
		return
	}

	user, err := u.UserService.GetUser(&userRequests.GetUserRequest{UUID: request.AssigneeUUID})
	if err != nil {
		if errors.Is(err, appErrors.UserNotFound) {
			responses.WriteResponse(w, responses.NewNotFoundRequest(), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}

	_, err = u.TaskService.CreateTask(request, user)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, responses.NewSuccessResponse(), u.logger)
}

// Get return user by UUID
func (u *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetTaskRequest(r)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	err = u.validator.Validate(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors(u.validator.GetErrors(err)), u.logger)
		return
	}
	task, err := u.TaskService.GetTask(request)
	if err != nil {
		if errors.Is(err, appErrors.TaskNotFound) {
			responses.WriteResponse(w, responses.NewNotFoundRequest(), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, taskResponses.NewGetTaskResponse(task), u.logger)
}
// Delete deletes task
func (u *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteTaskRequest(r)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	err = u.validator.Validate(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors(u.validator.GetErrors(err)), u.logger)
		return
	}
	err = u.TaskService.DeleteTask(request)
	if err != nil {
		if errors.Is(err, appErrors.TaskNotFound) {
			responses.WriteResponse(w, responses.NewNotFoundRequest(), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, responses.NewSuccessResponse(), u.logger)
}
// Search returns tasks by given condition
func (u *TaskHandler) Search(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSearchTaskRequest(r)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	err = u.validator.Validate(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors(u.validator.GetErrors(err)), u.logger)
		return
	}
	tasks, err := u.TaskService.SearchTasks(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, taskResponses.NewSearchTasksResponse(tasks), u.logger)
}

// SetStatus sets status for the task
func (u *TaskHandler) SetStatus(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSetTaskStatus(r)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	err = u.validator.Validate(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors(u.validator.GetErrors(err)), u.logger)
		return
	}
	err = u.TaskService.SetTaskStatus(request)
	if err != nil {
		if errors.Is(err, appErrors.UserAlreadyHaveActiveTask) {
			responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors([]string{appErrors.UserAlreadyHaveActiveTask.Error()}), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, responses.NewSuccessResponse(), u.logger)
}

// SetReminder sets reminder on the task
func (u *TaskHandler) SetReminder(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSetReminderRequest(r)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	err = u.validator.Validate(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors(u.validator.GetErrors(err)), u.logger)
		return
	}
	err = u.TaskService.SetReminder(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, responses.NewSuccessResponse(), u.logger)
}

// AddWorklog adds worklog to the task
func (u *TaskHandler) AddWorklog(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewAddWorklogRequest(r)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	err = u.validator.Validate(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors(u.validator.GetErrors(err)), u.logger)
		return
	}
	err = u.TaskService.AddWorklog(request)
	if err != nil {
		if errors.Is(err, appErrors.TaskNotFound) {
			responses.WriteResponse(w, responses.NewNotFoundRequest(), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, responses.NewSuccessResponse(), u.logger)
}
