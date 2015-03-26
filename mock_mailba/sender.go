package mock_mailba

import "github.com/plimble/mailba"
import "github.com/stretchr/testify/mock"

type MockSender struct {
	mock.Mock
}

func NewMockSender() *MockSender {
	return &MockSender{}
}

func (m *MockSender) Send(mail *mailba.Mail) error {
	ret := m.Called(mail)

	r0 := ret.Error(0)

	return r0
}
