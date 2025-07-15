package mailer

import (
	"bytes"
	"fmt"
	"log"
	"mailer-service/config"
	"mailer-service/internal/model"
	"mailer-service/internal/service"

	"github.com/wneessen/go-mail"
)

type Mailer struct {	
	templateService *service.TemplateService
	config          *config.Config
}

func NewMailer(templateService *service.TemplateService, config *config.Config) *Mailer {
	return &Mailer{
		templateService: templateService,
		config:		  config,
	}
}

func (m *Mailer) SendEmail(request *model.Request) error {
	body, err := m.templateService.Render(request.TemplateName, request)
	if err != nil {		
		return fmt.Errorf("failed rendering template: %w", err)
		
	}
	message := mail.NewMsg()
	if err := message.From("info@zenith.ovh"); err != nil {
		return fmt.Errorf("failed to set From address: %s", err)
		
	}
	if err := message.To(request.MailTo); err != nil {
		return fmt.Errorf("failed to set To address: %s", err)
		
	}
	if request.MailCc != nil {
		if err := message.Cc(*request.MailCc); err != nil {
			return fmt.Errorf("failed to set Cc address: %s", err)
			
		}
	}		
	if request.Attachments != nil {
		for _, att := range *request.Attachments {
			err := message.AttachReader(att.Filename, bytes.NewReader(att.Content),  mail.WithFileContentType("application/pdf"))
			if err != nil {
				return err
			}
		}
	}
	message.Subject(request.Subject)
	message.SetBodyString(mail.TypeTextHTML, body)
	client, err := mail.NewClient(m.config.SMTPServer, mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(m.config.SMTPUser), mail.WithPassword(m.config.SMTPPassword))
	if err != nil {
		return fmt.Errorf("failed to create mail client: %s", err)		
	}
	log.Printf("Message has %d attachments", len(message.GetAttachments()))
	if err := client.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send mail: %s", err)				
	}
	return nil
}