package messaging

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.nanomsg.org/mangos/v3"
)

type RequesterTest struct {
	suite.Suite
	assert *assert.Assertions
	sub    *Requester
	log    *mockLogger
	socket *mockSocket
}

func (t *RequesterTest) SetupTest() {
	t.assert = assert.New(t.T())
	t.log = &mockLogger{}
	t.socket = &mockSocket{}
	t.sub = NewRequester().WithLogger(t.log)
	t.assert.NotNil(t.sub)
}

func TestRequester(t *testing.T) {
	suite.Run(t, new(RequesterTest))
}

func (t *RequesterTest) TestSendRequestNewSocketFail() {
	newReqSocket = func() (mangos.Socket, error) {
		return nil, errors.New("anError")
	}
	t.log.On("Debug", "Failed to create request socket").Return()
	msg, err := t.sub.SendRequest([]byte{}, "test")
	t.assert.Equal(msg, []byte(nil))
	t.assert.Equal(err, errors.New("anError"))
}

func (t *RequesterTest) TestSendRequestDialFail() {
	newReqSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Dial", "ipc://test_point_replier").Return(errors.New("anError"))
	t.log.On("Debug", "Failed to dial on request socket").Return()
	msg, err := t.sub.SendRequest([]byte{}, "test")
	t.assert.Equal(msg, []byte(nil))
	t.assert.Equal(err, errors.New("anError"))
}

func (t *RequesterTest) TestSendRequestSendFail() {
	newReqSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Dial", "ipc://test_point_replier").Return(nil)
	t.socket.On("Send", []byte{}).Return(errors.New("anError"))
	t.log.On("Debug", "Failed to send request").Return()
	msg, err := t.sub.SendRequest([]byte{}, "test")
	t.assert.Equal(msg, []byte(nil))
	t.assert.Equal(err, errors.New("anError"))
}

func (t *RequesterTest) TestSendRequestRecvFail() {
	newReqSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Dial", "ipc://test_point_replier").Return(nil)
	t.socket.On("Send", []byte{}).Return(nil)
	t.socket.On("Recv").Return([]byte(nil), errors.New("anError"))
	t.log.On("Debug", "Failed to receive reply").Return()
	msg, err := t.sub.SendRequest([]byte{}, "test")
	t.assert.Equal(msg, []byte(nil))
	t.assert.Equal(err, errors.New("anError"))
}

func (t *RequesterTest) TestSendRequestCloseFail() {
	newReqSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Dial", "ipc://test_point_replier").Return(nil)
	t.socket.On("Send", []byte{}).Return(nil)
	t.socket.On("Recv").Return([]byte{}, nil)
	t.socket.On("Close").Return(errors.New("anError"))
	t.log.On("Debug", "Failed to close request socket").Return()
	msg, err := t.sub.SendRequest([]byte{}, "test")
	t.assert.Equal(msg, []byte{})
	t.assert.Equal(err, nil)
}

func (t *RequesterTest) TestSendRequestSuccess() {
	newReqSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Dial", "ipc://test_point_replier").Return(nil)
	t.socket.On("Send", []byte{}).Return(nil)
	t.socket.On("Recv").Return([]byte{}, nil)
	t.socket.On("Close").Return(nil)
	msg, err := t.sub.SendRequest([]byte{}, "test")
	t.assert.Equal(msg, []byte{})
	t.assert.Equal(err, nil)
}
