package point

import "github.com/stretchr/testify/mock"

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Debug(msg string) {
	m.Called(msg)
}

type mockPoint struct {
	mock.Mock
}

func (m *mockPoint) Execute() {
	m.Called()
}

func (m *mockPoint) Resume() {
	m.Called()
}
