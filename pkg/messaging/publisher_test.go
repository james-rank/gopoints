package messaging

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.nanomsg.org/mangos/v3"
)

type PublisherTest struct {
	suite.Suite
	assert *assert.Assertions
	sub    *Publisher
	log    *mockLogger
	socket *mockSocket
}

func (t *PublisherTest) SetupTest() {
	t.assert = assert.New(t.T())
	t.log = &mockLogger{}
	t.socket = &mockSocket{}
	t.sub = NewPublisher().WithLogger(t.log)
	t.assert.NotNil(t.sub)
}

func TestPublisher(t *testing.T) {
	suite.Run(t, new(PublisherTest))
}

func (t *PublisherTest) TestStartAlreadyStarted() {
	t.log.On("Debug", "Publisher already started").Return()
	t.sub.started.Store(true)
	t.assert.Equal(ErrAlreadyStarted, t.sub.Start("test"))
}

func (t *PublisherTest) TestStartNewSocketFail() {
	newPubSocket = func() (mangos.Socket, error) {
		return nil, errors.New("anError")
	}
	t.log.On("Debug", "Failed to create publisher socket").Return()
	t.assert.Equal(t.sub.Start("test"), errors.New("anError"))
}

func (t *PublisherTest) TestStartListenFail() {
	newPubSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Listen", "ipc://test_point_publish").Return(errors.New("anError"))
	t.log.On("Debug", "Failed to listen on publisher socket").Return()
	t.assert.Equal(t.sub.Start("test"), errors.New("anError"))
}

func (t *PublisherTest) TestStartSuccess() {
	newPubSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Listen", "ipc://test_point_publish").Return(nil)
	t.assert.Nil(t.sub.Start("test"))
}

func (t *PublisherTest) TestStopAlreadyStopped() {
	t.log.On("Debug", "Publisher already stopped").Return()
	t.assert.Equal(t.sub.Stop(), ErrAlreadyStopped)
}

func (t *PublisherTest) TestStopCloseFail() {
	t.sub.started.Store(true)
	t.sub.socket = t.socket
	t.socket.On("Close").Return(errors.New("anError"))
	t.log.On("Debug", "Failed to close publisher socket").Return()
	t.assert.Equal(t.sub.Stop(), errors.New("anError"))
}

func (t *PublisherTest) TestStopSuccess() {
	t.sub.started.Store(true)
	t.sub.socket = t.socket
	t.socket.On("Close").Return(nil)
	t.assert.Nil(t.sub.Stop())
}

func (t *PublisherTest) TestPublisherNotStarted() {
	t.sub.started.Store(false)
	t.log.On("Debug", "Publisher not started").Return()
	t.assert.Equal(t.sub.Publish([]byte{}), ErrNotStarted)
}

func (t *PublisherTest) TestPublishFail() {
	t.sub.started.Store(true)
	t.sub.socket = t.socket
	t.socket.On("Send", []byte{}).Return(errors.New("anError"))
	t.log.On("Debug", "Failed to send message").Return()
	t.assert.Equal(t.sub.Publish([]byte{}), errors.New("anError"))
}

func (t *PublisherTest) TestPublishSuccess() {
	t.sub.started.Store(true)
	t.sub.socket = t.socket
	t.socket.On("Send", []byte{}).Return(nil)
	t.assert.Nil(t.sub.Publish([]byte{}))
}
