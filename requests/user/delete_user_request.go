package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DeleteUserRequest struct {
	UUID string `validate:"required",json:"uuid"`
}

func NewDeleteUserRequest(r *http.Request) (*DeleteUserRequest, error) {
	request := &DeleteUserRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
