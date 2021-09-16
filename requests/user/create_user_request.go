package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Email     string `validate:"required",json:"email"`
	FirstName string `validate:"required",json:"firstName"`
	LastName  string `validate:"required",json:"lastName"`
}

func NewCreateUserRequest(r *http.Request) (*CreateUserRequest, error) {
	request := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
