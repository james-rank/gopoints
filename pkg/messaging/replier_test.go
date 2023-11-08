package messaging

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.nanomsg.org/mangos/v3"
)

type ReplierTest struct {
	suite.Suite
	assert *assert.Assertions
	sub    *Replier
	log    *mockLogger
	socket *mockSocket
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func (t *ReplierTest) SetupTest() {
	t.assert = assert.New(t.T())
	t.log = &mockLogger{}
	t.socket = &mockSocket{}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	t.wg = &sync.WaitGroup{}
	t.sub = NewReplier().WithLogger(t.log)
	t.assert.NotNil(t.sub)
}

func TestReplier(t *testing.T) {
	suite.Run(t, new(ReplierTest))
}

func (t *ReplierTest) TestRunNewSocketFail() {
	newRepSocket = func() (mangos.Socket, error) {
		return nil, errors.New("anError")
	}
	t.log.On("Debug", "Failed to create reply socket").Return()
	t.wg.Add(1)
	t.assert.Equal(t.sub.Run(t.ctx, t.wg, "", func(b []byte) []byte { return nil }), errors.New("anError"))
}

func (t *ReplierTest) TestRunListenFail() {
	newRepSocket = func() (mangos.Socket, error) {
		return t.socket, nil
	}
	t.socket.On("Listen", "ipc://_point_replier").Return(errors.New("anError"))
	t.log.On("Debug", "Failed to listen on reply socket").Return()
	t.wg.Add(1)
	t.assert.Equal(t.sub.Run(t.ctx, t.wg, "", func(b []byte) []byte { return nil }), errors.New("anError"))
}
