package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateTaskRequest struct {
	AssigneeUUID string `validate:"required",json:"assigneeUUID"`
	Title        string `validate:"required",json:"title"`
}

func NewCreateTaskRequest(r *http.Request) (*CreateTaskRequest, error) {
	request := &CreateTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
