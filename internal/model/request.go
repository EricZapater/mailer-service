package model

type Attachment struct {
	Filename string
	Content  []byte
}

type Request struct {
	MailTo       string  `json:"mail_to"`
	MailCc       *string `json:"mail_cc"`
	Body         string  `json:"body"`
	TemplateName string  `json:"template_name"`
	Subject      string  `json:"subject"`
	Footer       string  `json:"footer"`
	Attachments  *[]Attachment
}