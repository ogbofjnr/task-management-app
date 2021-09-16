package repositories

import (
	sql2 "database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/ogbofjnr/maze/constants"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/errors"
	"github.com/ogbofjnr/maze/pkg/db"
	requests "github.com/ogbofjnr/maze/requests/user"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Insert inserts user into the db
func (r *UserRepository) Insert(db db.DB, user *entities.User) error {
	_, err := db.NamedExec(`
	INSERT INTO users (uuid,email,first_name,last_name,status,created_at,updated_at)
	VALUES (:uuid, :email, :first_name, :last_name, :status, now(), now())`, user)

	if err != nil {
		return fmt.Errorf("error inserting user to db: %w", err)
	}

	return nil
}

// GetByUUID finds user by UUID
func (r *UserRepository) GetByUUID(db db.DB, uuid string) (*entities.User, error) {
	user := &entities.User{}
	sqStmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := sqStmt.Select("*").From("users").Where("status!=? AND uuid=?", constants.Deleted, uuid).ToSql()

	if err != nil {
		return nil, fmt.Errorf("error query sql: %w", err)
	}

	err = db.Get(user, sql, args...)

	if err == sql2.ErrNoRows {
		return nil, errors.UserNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("error selecting user from db: %w", err)
	}
	return user, nil
}

// GetByEmail finds user by email
func (r *UserRepository) GetByEmail(db db.DB, email string) (*entities.User, error) {
	user := &entities.User{}
	sqStmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := sqStmt.Select("*").From("users").Where("email=?", email).ToSql()

	if err != nil {
		return nil, fmt.Errorf("error query sql: %w", err)
	}

	err = db.Get(user, sql, args...)

	if err == sql2.ErrNoRows {
		return nil, errors.UserNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("error selecting user from db: %w", err)
	}

	return user, nil
}

// GetByID finds user by ID
func (r *UserRepository) GetByID(db db.DB, ID int) (*entities.User, error) {
	user := &entities.User{}
	sqStmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := sqStmt.Select("*").From("users").Where("ID=?", ID).ToSql()

	if err != nil {
		return nil, fmt.Errorf("error creating sql query : %w", err)
	}

	err = db.Get(user, sql, args...)

	if err == sql2.ErrNoRows {
		return nil, errors.UserNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("error selecting user from db: %w", err)
	}

	return user, nil
}

// Delete soft deletes  user
func (r *UserRepository) Delete(db db.DB, user *entities.User) error {
	_, err := db.Exec(`UPDATE users SET status=$1, deleted_at=now() WHERE id=$2`, constants.Deleted, user.ID)

	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}

// Search finds users by given condition
func (r *UserRepository) Search(db db.DB, request *requests.SearchUserRequest) (*[]entities.User, error) {
	users := &[]entities.User{}
	sqStmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := sqStmt.Select("*").From("users")

	if request.Email != nil {
		builder = builder.Where("email=?", request.Email)
	}

	if request.FirstName != nil {
		builder = builder.Where("first_name=?", request.FirstName)
	}

	if request.LastName != nil {
		builder = builder.Where("last_name=?", request.LastName)
	}

	if request.Status != nil {
		builder = builder.Where("status=?", request.Status)
	}

	sql, args, err := builder.ToSql()

	if err != nil {
		return nil, fmt.Errorf("error query sql: %w", err)
	}

	err = db.Select(users, sql, args...)

	if err != nil {
		return nil, fmt.Errorf("error selecting user from db: %w", err)
	}
	return users, nil
}
