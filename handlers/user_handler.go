package handlers

import (
	"errors"
	appErrors "github.com/ogbofjnr/maze/errors"
	validator2 "github.com/ogbofjnr/maze/pkg/validator"
	requests "github.com/ogbofjnr/maze/requests/user"
	"github.com/ogbofjnr/maze/responses"
	userResponses "github.com/ogbofjnr/maze/responses/user"
	"github.com/ogbofjnr/maze/services"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	validator   validator2.Validator
	UserService *services.UserService
	logger      *zap.Logger
}

func NewUserHandler(validator validator2.Validator, userService *services.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		validator:   validator,
		UserService: userService,
		logger:      logger,
	}
}

// Create creates the user
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateUserRequest(r)
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
	_, err = u.UserService.CreateUser(request)
	if err != nil {
		if errors.Is(err, appErrors.UserAlreadyExists) {
			responses.WriteResponse(w, responses.NewBadRequestResponse().SetErrors([]string{err.Error()}), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, responses.NewSuccessResponse(), u.logger)
}

// Get gets user by UUID
func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetUserRequest(r)
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
	user, err := u.UserService.GetUser(request)
	if err != nil {
		if errors.Is(err, appErrors.UserNotFound) {
			responses.WriteResponse(w, responses.NewNotFoundRequest(), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, userResponses.NewGetUserResponse(user), u.logger)
}

// Delete deletes user
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteUserRequest(r)
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
	err = u.UserService.DeleteUser(request)
	if err != nil {
		if errors.Is(err, appErrors.UserNotFound) {
			responses.WriteResponse(w, responses.NewNotFoundRequest(), u.logger)
			return
		}
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, responses.NewSuccessResponse(), u.logger)
}

// Search returns users by given condition
func (u *UserHandler) Search(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSearchUserRequest(r)
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
	users, err := u.UserService.SearchUsers(request)
	if err != nil {
		u.logger.Error(err.Error())
		responses.WriteResponse(w, responses.NewInternalServerErrorResponse(), u.logger)
		return
	}
	responses.WriteResponse(w, userResponses.NewSearchUsersResponse(users), u.logger)
}
