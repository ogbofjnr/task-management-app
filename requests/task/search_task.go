package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SearchTaskRequest struct {
	Title  *string `validate:"",json:"title,omitempty"`
	Status *string `validate:"required",json:"status"`
}

func NewSearchTaskRequest(r *http.Request) (*SearchTaskRequest, error) {
	request := &SearchTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
