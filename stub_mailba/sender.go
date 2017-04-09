package stub_mailba

import (
	"github.com/plimble/errors"
	"github.com/devlist/mailba"
)

type StubSender struct{}

func NewStubSender() *StubSender {
	return &StubSender{}
}

func (s *StubSender) Send(mail *mailba.Mail) error {
	if mail == nil {
		return errors.InternalError("no mail object")
	}

	if mail.Subject == "error" {
		return errors.InternalError("send mail failed")
	}

	return nil
}
