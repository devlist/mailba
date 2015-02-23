// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/plimble/mailba (interfaces: Sender)

package mock_mailba

import (
	mailba "github.com/plimble/mailba"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *_MockSenderRecorder
}

// Recorder for MockSender (not exported)
type _MockSenderRecorder struct {
	mock *MockSender
}

func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &_MockSenderRecorder{mock}
	return mock
}

func (_m *MockSender) EXPECT() *_MockSenderRecorder {
	return _m.recorder
}

func (_m *MockSender) Send(_param0 *mailba.Mail, _param1 *mailba.Config) error {
	ret := _m.ctrl.Call(_m, "Send", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSenderRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Send", arg0, arg1)
}
