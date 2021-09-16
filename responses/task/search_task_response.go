package user_responses

import (
	"encoding/json"
	"fmt"
	"github.com/ogbofjnr/maze/database/entities"
	"github.com/ogbofjnr/maze/responses"
	"net/http"
)

type Task struct {
	UUID   string `json:"uuid"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type SearchTasksResponse struct {
	responses.Response
	Tasks []Task
}

func (a *SearchTasksResponse) ToString() ([]byte, error) {
	j, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling response :%w", err)
	}
	return j, nil
}

func NewSearchTasksResponse(tasks *[]entities.Task) *SearchTasksResponse {
	r := &SearchTasksResponse{}
	r.StatusCode = http.StatusOK
	for _, entity := range *tasks {
		t := Task{}
		t.UUID = entity.UUID
		t.Title = entity.Title
		t.Status = entity.Status
		r.Tasks = append(r.Tasks, t)
	}
	return r
}
