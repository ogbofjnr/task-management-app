package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DeleteTaskRequest struct {
	UUID string `validate:"required",json:"uuid"`
}

func NewDeleteTaskRequest(r *http.Request) (*DeleteTaskRequest, error) {
	request := &DeleteTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
