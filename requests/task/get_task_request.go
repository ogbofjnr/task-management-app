package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetTaskRequest struct {
	UUID string `validate:"required",json:"uuid"`
}

func NewGetTaskRequest(r *http.Request) (*GetTaskRequest, error) {
	request := &GetTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
