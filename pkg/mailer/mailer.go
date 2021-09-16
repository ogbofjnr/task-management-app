package mailer

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	sendgridMail "github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.uber.org/zap"
)

const templateFolder = "email_templates"

type Mail interface {
	Render() (string, error)
	GetFrom() *sendgridMail.Email
	GetTo() *sendgridMail.Email
	GetSubject() string
}

type Mailer struct {
	client         *sendgrid.Client
	templateFolder string
	logger         *zap.Logger
}

func NewMailer(client *sendgrid.Client, logger *zap.Logger) *Mailer {
	return &Mailer{
		client:         client,
		templateFolder: templateFolder,
		logger:         logger,
	}
}

func (m *Mailer) Send(mail Mail) error {

	mailData, err := mail.Render()
	if err != nil {
		return fmt.Errorf("error sending email :%w", err)
	}

	message := sendgridMail.NewSingleEmail(mail.GetFrom(), mail.GetSubject(), mail.GetTo(), mailData, mailData)

	_, err = m.client.Send(message)
	if err != nil {
		return fmt.Errorf("error sending email :%w", err)
	}

	return nil
}
