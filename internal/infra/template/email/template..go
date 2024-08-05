package email_template

import (
	"bytes"
	"html/template"
	"path"
)

const (
	PotentialCompetitorsTemplate = "potential_competitors.template.html"
)

func ParseTemplate(templateFileName string, data interface{}) (*string, error) {
	templatePath := path.Join("internal", "infra", "template", "email", templateFileName)

	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	htmlBuffer := new(bytes.Buffer)
	if err = template.Execute(htmlBuffer, data); err != nil {
		return nil, err
	}

	html := htmlBuffer.String()

	return &html, nil
}
