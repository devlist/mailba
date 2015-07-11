package stub_mailba

import (
	"errors"
	"github.com/plimble/mailba"
)

type StubSender struct{}

func NewStubSender() *StubSender {
	return &StubSender{}
}

func (s *StubSender) Send(mail *mailba.Mail) error {
	if mail == nil {
		return errors.New("no mail object")
	}

	if mail.Subject == "error" {
		return errors.New("send mail failed")
	}

	return nil
}
