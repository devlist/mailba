package mailba

import (
	"github.com/mattbaird/gochimp"
	"github.com/plimble/errs"
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
		Subject:   mail.subject,
		FromEmail: mail.from.email,
		FromName:  mail.from.name,
	}

	msg.Html = mail.content

	for _, v := range mail.to {
		msg.AddRecipients(gochimp.Recipient{Email: v.email, Name: v.name, Type: "to"})
	}

	for _, v := range mail.cc {
		msg.AddRecipients(gochimp.Recipient{Email: v.email, Name: v.name, Type: "cc"})
	}

	for _, v := range mail.bcc {
		msg.AddRecipients(gochimp.Recipient{Email: v.email, Name: v.name, Type: "bcc"})
	}

	for _, v := range mail.files {
		msg.AddAttachments(gochimp.Attachment{Content: v.content, Name: v.name, Type: v.mime})
	}

	return msg
}

func (m *Mandrill) Send(mail *Mail, config *Config) error {
	var err error

	// msg.AddMergeVar(gochimp.MergeVars{
	// 	Recipient: email,
	// 	Vars: []gochimp.Var{
	// 		gochimp.Var{"name", name},
	// 		gochimp.Var{"password", newPwd},
	// 	},
	// })

	if mail.template != "" {
		contentVar := []gochimp.Var{}
		msg := m.mapMsg(mail)
		_, err = m.api.MessageSendTemplate(`ams-resetpassword`, contentVar, msg, true)
	} else {
		msg := m.mapMsg(mail)
		_, err = m.api.MessageSend(msg, true)
	}

	if err != nil {
		return errs.NewErrors(err.Error())
	}

	return nil
}
