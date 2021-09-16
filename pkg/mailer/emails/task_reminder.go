package emails

import (
	"bytes"
	"fmt"
	"github.com/ogbofjnr/maze/constants"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"os"
)

const htmlTemplate = "task_reminder_template.html"
const taskReminderSubject = "task reminder"
const fromAddress = "notifications@domain.com"
const fromName = "Domain support"

type TaskReminder struct {
	Mail
	UserName  string
	TaskTitle string
}

func NewTaskReminder(to string, userName string, taskTitle string) *TaskReminder {
	return &TaskReminder{
		Mail: Mail{
			From:     mail.NewEmail(fromName, fromAddress),
			To:       mail.NewEmail(userName, to),
			Subject:  taskReminderSubject,
			Template: htmlTemplate,
		},
		UserName:  userName,
		TaskTitle: taskTitle,
	}
}

func (t *TaskReminder) Render() (string, error) {
	dir, _ := os.Getwd()
	var tpl = template.Must(template.ParseGlob(dir + constants.TemplateDir + t.Template))

	data := struct {
		UserName  string
		TaskTitle string
	}{
		t.UserName,
		t.TaskTitle,
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("error rendering email template :%w", err)
	}

	return buf.String(), nil
}
