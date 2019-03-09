package email

import (
	"bytes"
	"html/template"
	"net/smtp"
)

// Email settings
type Email struct {
	config *Config
	auth   smtp.Auth
	tpl    *template.Template
	body   string
}

// New returns the type email with default settings
func New(cfg *Config) *Email {
	auth := smtp.PlainAuth("", cfg.From, cfg.Password, cfg.Host)
	t := template.New("Email")

	return &Email{
		config: cfg,
		auth:   auth,
		tpl:    t,
	}
}

// Send the e-mail
func (e *Email) Send(subject string, data interface{}) (bool, error) {

	err := e.parseData(data)
	if err != nil {
		return false, err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "!\n"
	msg := []byte(subject + mime + "\n" + e.body)
	addr := e.config.Host + ":" + e.config.Port

	if err := smtp.SendMail(addr, e.auth, e.config.From, []string{e.config.To}, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (e *Email) parseData(data interface{}) error {

	t, err := e.tpl.Parse(emailTpl)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return err
	}

	e.body = buf.String()
	return nil
}
