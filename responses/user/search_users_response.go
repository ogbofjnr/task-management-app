package user_responses

import (
	"encoding/json"
	"fmt"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/responses"
	"net/http"
)

type User struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Status    string `json:"status"`
}

type SearchUsersResponse struct {
	responses.Response
	Users []User
}

func (a *SearchUsersResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func NewSearchUsersResponse(users *[]entities.User) *SearchUsersResponse {
	r := &SearchUsersResponse{}
	r.StatusCode = http.StatusOK
	for _, entity := range *users {
		u := User{}
		u.UUID = entity.UUID
		u.Email = entity.Email
		u.FirstName = entity.FirstName
		u.LastName = entity.LastName
		u.Status = entity.Status
		r.Users = append(r.Users, u)
	}
	return r
}
