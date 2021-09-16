package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AddWorklogRequest struct {
	TaskUUID   string `validate:"required",json:"taskUUID"`
	LoggedTime int    `validate:"required",json:"loggedTime"`
}

func NewAddWorklogRequest(r *http.Request) (*AddWorklogRequest, error) {
	request := &AddWorklogRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
