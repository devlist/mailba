package mailba

import (
	"github.com/mattbaird/gochimp"
	"github.com/plimble/utils/errors2"
)

type Mandrill struct {
	api *gochimp.MandrillAPI
}

func NewMandrill(apiKey string) (*Mandrill, error) {
	api, err := gochimp.NewMandrill(apiKey)
	if err != nil {
		return nil, err
	}

	return &Mandrill{api}, nil
}

func (m *Mandrill) mapMsg(mail *Mail) gochimp.Message {
	msg := gochimp.Message{
		Subject:   mail.Subject,
		FromEmail: mail.From.Email,
		FromName:  mail.From.Name,
	}

	msg.Html = mail.Content

	for _, v := range mail.To {
		msg.AddRecipients(gochimp.Recipient{Email: v.Email, Name: v.Name, Type: "to"})
	}

	for _, v := range mail.CC {
		msg.AddRecipients(gochimp.Recipient{Email: v.Email, Name: v.Name, Type: "cc"})
	}

	for _, v := range mail.BCC {
		msg.AddRecipients(gochimp.Recipient{Email: v.Email, Name: v.Name, Type: "bcc"})
	}

	for _, v := range mail.Files {
		msg.AddAttachments(gochimp.Attachment{Content: v.Content, Name: v.Name, Type: v.Mime})
	}

	return msg
}

func (m *Mandrill) Send(mail *Mail, config *Config) error {
	var err error

	msg := m.mapMsg(mail)

	for email, vars := range mail.Mergevars {
		msg.AddMergeVar(gochimp.MergeVars{
			Recipient: email,
			Vars:      vars,
		})
	}

	msg.GlobalMergeVars = mail.Globalvars

	if mail.Template != "" {
		contentVar := []gochimp.Var{}
		_, err = m.api.MessageSendTemplate(mail.Template, contentVar, msg, true)
	} else {
		_, err = m.api.MessageSend(msg, true)
	}

	if err != nil {
		return errors2.NewInternal(err.Error())
	}

	return nil
}
