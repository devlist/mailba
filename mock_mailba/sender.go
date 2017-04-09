package mock_mailba

import "github.com/devlist/mailba"
import "github.com/stretchr/testify/mock"

type Sender struct {
	mock.Mock
}

func (_m *Sender) Send(mail *mailba.Mail) error {
	ret := _m.Called(mail)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mailba.Mail) error); ok {
		r0 = rf(mail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
