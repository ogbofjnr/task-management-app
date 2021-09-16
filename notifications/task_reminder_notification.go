package notifications

import (
	"encoding/json"
	"fmt"
)

const notificationType = "task_reminder"

type TaskReminderNotification struct {
	Type      string `json:"type"`
	TaskTitle string `json:"taskTitle"`
	TaskUUID  string `json:"taskUUID"`
}

func NewTaskReminderNotification(taskTitle, taskUUID string) *TaskReminderNotification {
	return &TaskReminderNotification{
		Type:      notificationType,
		TaskTitle: taskTitle,
		TaskUUID:  taskUUID,
	}
}

// GetMessage returns notification payload
func (t *TaskReminderNotification) GetMessage() ([]byte, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling notification :%w", err)
	}
	return data, nil
}
