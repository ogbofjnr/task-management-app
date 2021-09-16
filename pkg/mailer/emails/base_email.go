package emails

import "github.com/sendgrid/sendgrid-go/helpers/mail"

type Mail struct {
	From     *mail.Email
	To       *mail.Email
	Subject  string
	Template string
}

func (t *Mail) GetFrom() *mail.Email {
	return t.From
}

func (t *Mail) GetTo() *mail.Email {
	return t.To
}

func (t *Mail) GetSubject() string {
	return t.Subject
}
