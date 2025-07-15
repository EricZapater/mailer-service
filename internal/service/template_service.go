package service

import (
	"bytes"
	"fmt"
	"html/template"
)

type TemplateService struct{}

func NewTemplateService() *TemplateService {
	return &TemplateService{}
}

func (ts *TemplateService) Render(templateName string, data any) (string, error) {
	tmplPath := fmt.Sprintf("templates/%s.html", templateName)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
