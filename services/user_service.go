package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ogbofjnr/maze/constants"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/database/repositories"
	appErrors "github.com/ogbofjnr/maze/errors"
	requests "github.com/ogbofjnr/maze/requests/user"
)

type UserService struct {
	userRepository *repositories.UserRepository
	db             *sqlx.DB
}

func NewUserService(UserRepository *repositories.UserRepository, db *sqlx.DB) *UserService {
	return &UserService{
		userRepository: UserRepository,
		db:             db,
	}
}

func (s *UserService) CreateUser(request *requests.CreateUserRequest) (*entities.User, error) {
	existingUser, err := s.userRepository.GetByEmail(s.db, request.Email)

	if err != nil && !errors.Is(err, appErrors.UserNotFound) {
		return nil, fmt.Errorf("error getting the user: %w", err)
	}

	if existingUser != nil {
		return nil, appErrors.UserAlreadyExists
	}

	user := &entities.User{}
	user.UUID = uuid.New().String()
	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Status = constants.Active
	err = s.userRepository.Insert(s.db, user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	return user, nil
}

func (s *UserService) GetUser(request *requests.GetUserRequest) (*entities.User, error) {
	user, err := s.userRepository.GetByUUID(s.db, request.UUID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return user, nil
}

func (s *UserService) DeleteUser(request *requests.DeleteUserRequest) error {
	user, err := s.userRepository.GetByUUID(s.db, request.UUID)
	if err != nil {
		return fmt.Errorf("error retrieving user: %w", err)
	}
	err = s.userRepository.Delete(s.db, user)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (s *UserService) SearchUsers(request *requests.SearchUserRequest) (*[]entities.User, error) {
	users, err := s.userRepository.Search(s.db, request)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return users, nil
}
