package model

type Attachment struct {
	Filename    string `json:"filename" example:"document.pdf"`
	Content     []byte `json:"content" swaggerignore:"true"`
	ContentType string `json:"content_type" example:"application/pdf"`
}

type Request struct {
	MailTo       string        `json:"mail_to" example:"user@example.com"`
	MailCc       *string       `json:"mail_cc,omitempty" example:"cc@example.com"`
	Subject      string        `json:"subject" example:"Nova factura"`
	Body         string        `json:"body" example:"Contingut del email"`
	TemplateName string        `json:"template_name" example:"invoice"`
	Footer       string        `json:"footer" example:"Peu del email"`
	Attachments  *[]Attachment `json:"attachments,omitempty"`
}