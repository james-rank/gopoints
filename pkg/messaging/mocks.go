package messaging

import (
	"github.com/stretchr/testify/mock"
	"go.nanomsg.org/mangos/v3"
)

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Debug(arg0 string) {
	m.Called(arg0)
}

type mockSocket struct {
	mangos.Socket
	mock.Mock
}

func (m *mockSocket) Close() error {
	ret := m.Called()

	r0 := ret.Error(0)

	return r0
}

func (m *mockSocket) Dial(arg0 string) error {
	ret := m.Called(arg0)

	r0 := ret.Error(0)

	return r0
}

func (m *mockSocket) Listen(arg0 string) error {
	ret := m.Called(arg0)

	r0 := ret.Error(0)

	return r0
}

func (m *mockSocket) Send(arg0 []byte) error {
	ret := m.Called(arg0)

	r0 := ret.Error(0)

	return r0
}

func (m *mockSocket) SetOption(arg0 string, arg1 interface{}) error {
	ret := m.Called(arg0, arg1)

	r0 := ret.Error(0)

	return r0
}

func (m *mockSocket) Recv() ([]byte, error) {
	ret := m.Called()

	r0 := ret.Get(0).([]byte)
	r1 := ret.Error(1)

	return r0, r1
}