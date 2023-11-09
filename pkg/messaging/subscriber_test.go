package messaging

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SubscriberTest struct {
	suite.Suite
	assert *assert.Assertions
	sub    *Subscriber
	log    *mockLogger
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
	socket *mockSocket
}

func (t *SubscriberTest) SetupTest() {
	t.assert = assert.New(t.T())
	t.log = &mockLogger{}
	t.socket = &mockSocket{}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	t.wg = &sync.WaitGroup{}
	t.sub = NewSubscriber().WithLogger(t.log)
	t.assert.NotNil(t.sub)
}

func TestSubscriber(t *testing.T) {
	suite.Run(t, new(SubscriberTest))
}
