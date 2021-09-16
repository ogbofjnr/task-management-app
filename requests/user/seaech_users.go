package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SearchUserRequest struct {
	Email     *string `validate:"",json:"email,omitempty"`
	FirstName *string `validate:"",json:"firstName,omitempty"`
	LastName  *string `validate:"",json:"lastName,omitempty"`
	Status    *string `validate:"",json:"status,omitempty"`
}

func NewSearchUserRequest(r *http.Request) (*SearchUserRequest, error) {
	request := &SearchUserRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
