package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetUserRequest struct {
	UUID string `validate:"required",json:"uuid"`
}

func NewGetUserRequest(r *http.Request) (*GetUserRequest, error) {
	request := &GetUserRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
