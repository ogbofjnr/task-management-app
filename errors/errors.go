package errors

import "errors"

var UserAlreadyExists = errors.New("user already exists")
var UserNotFound = errors.New("user not found")
var TaskNotFound = errors.New("task not found")
var UserAlreadyHaveActiveTask = errors.New("user already has active task")
