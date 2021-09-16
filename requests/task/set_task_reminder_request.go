package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SetReminderRequest struct {
	TaskUUID     string    `validate:"required",json:"taskUUID"`
	ReminderTime time.Time `validate:"required",json:"reminderTime"`
	TaskStatus   string    `validate:"required",json:"taskStatus"`
}

func NewSetReminderRequest(r *http.Request) (*SetReminderRequest, error) {
	request := &SetReminderRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}
	return request, nil
}
