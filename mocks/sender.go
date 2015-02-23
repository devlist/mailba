package mocks

import "github.com/plimble/mailba"
import "github.com/stretchr/testify/mock"

type Sender struct {
	mock.Mock
}

func NewSender() *Sender {
	return &Sender{}
}

func (m *Sender) Send(mail *mailba.Mail, config *mailba.Config) error {
	ret := m.Called(mail, config)

	r0 := ret.Error(0)

	return r0
}
