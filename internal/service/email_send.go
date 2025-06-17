package service

import (
	"context"
	"e-marketing/internal/model"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailSendService interface {
	Send(ctx context.Context, sender model.Sender, recipient model.Recipient,
		subject, plainTextContent, htmlContent string) error
}

type sendgridService struct {
	client *sendgrid.Client
}

func (service *sendgridService) Send(ctx context.Context,
	sender model.Sender, recipient model.Recipient,
	subject, plainTextContent, htmlContent string) error {
	from := mail.NewEmail(sender.Name, sender.Email)
	to := mail.NewEmail(recipient.Name, recipient.Email)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	_, err := service.client.Send(message)
	return err
}

func NewEmailSendService(client *sendgrid.Client) EmailSendService {
	return &sendgridService{
		client: client,
	}
}
