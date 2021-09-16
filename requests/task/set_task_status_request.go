package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SetTaskStatus struct {
	TaskUUID string `validate:"required",json:"taskUUID"`
	Status   string `validate:"required",json:"status"`
}

func NewSetTaskStatus(r *http.Request) (*SetTaskStatus, error) {
	request := &SetTaskStatus{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
