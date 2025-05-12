package mailer

import (
	"context"
	"fmt"
	"log"
	"time"

	mailersend "github.com/mailersend/mailersend-go"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/model"
)

type Mailer struct {
	client *mailersend.Mailersend
}

func NewMailer(client *mailersend.Mailersend) *Mailer {
	return &Mailer{client: client}
}

func (m *Mailer) Send(ctx context.Context, customer model.Customer) error {

	ctxC, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := fmt.Sprintf("Welcome, %s!", customer.Email)
	text := "Greetings from the team, you got this message through MailerSend."
	html := "Greetings from the team, you got this message through MailerSend."

	from := mailersend.From{
		Name:  "AITU Coffee Shop",
		Email: "MS_14oy39@test-3m5jgrooznmgdpyo.mlsender.net",
	}

	recipients := []mailersend.Recipient{
		{
			Email: customer.Email,
		},
	}

	variables := []mailersend.Variables{
		{
			Email: customer.Email,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "foo",
					Value: "bar",
				},
			},
		},
	}

	tags := []string{"foo", "bar"}

	message := m.client.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetSubstitutions(variables)
	message.SetTags(tags)

	res, err := m.client.Email.Send(ctxC, message)
	if err != nil {
		log.Println("m.client.Email.Send:", err, "Status:", res.Status, "Customer:", customer.Email)
		log.Println("Try to send email again to: ", customer.Email)
		_, err = m.client.Email.Send(ctxC, message)
	}

	return err
}
