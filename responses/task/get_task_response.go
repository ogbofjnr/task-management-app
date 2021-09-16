package user_responses

import (
	"encoding/json"
	"fmt"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/responses"
	"net/http"
)

type GetTaskResponse struct {
	responses.Response
	UUID   string `json:"uuid"`
	Title  string `json:"email"`
	Status string `json:"status"`
}

func (a *GetTaskResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func NewGetTaskResponse(task *entities.Task) *GetTaskResponse {
	t := &GetTaskResponse{}
	t.StatusCode = http.StatusOK
	t.UUID = task.UUID
	t.Title = task.Title
	t.Status = task.Status
	return t
}
