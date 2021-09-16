package user_responses

import (
	"encoding/json"
	"fmt"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/responses"
	"net/http"
)

type GetUserResponse struct {
	responses.Response
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
}

func (a *GetUserResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func NewGetUserResponse(user *entities.User) *GetUserResponse {
	r := &GetUserResponse{}
	r.StatusCode = http.StatusOK
	r.UUID = user.UUID
	r.Email = user.Email
	r.FirstName = user.FirstName
	r.LastName = user.LastName
	r.Status = user.Status
	return r
}
