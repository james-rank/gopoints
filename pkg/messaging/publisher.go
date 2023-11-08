package messaging

import (
	"fmt"
	"sync/atomic"

	"github.com/james-rank/points/internal/pkg/log"
	"go.nanomsg.org/mangos/v3/protocol/pub"
	_ "go.nanomsg.org/mangos/v3/transport/all" // Required for transport registration.
)

var newPubSocket = pub.NewSocket

// Publisher is a struct that implements the Publisher interface.
type Publisher struct {
	logger  logger
	socket  socket
	started atomic.Bool
}

// NewPublisher creates a new Publisher.
func NewPublisher() *Publisher {
	return &Publisher{
		logger:  log.NewLogger(),
		started: atomic.Bool{},
	}
}

// WithLogger sets the logger for the Publisher.
func (p *Publisher) WithLogger(logger logger) *Publisher {
	p.logger = logger
	return p
}

// Start starts the Publisher.
func (p *Publisher) Start(uid string) (err error) {
	if p.started.Load() {
		p.logger.Debug("Publisher already started")
		return ErrAlreadyStarted
	}

	p.started.Store(true)

	p.socket, err = newPubSocket()
	if err != nil {
		p.logger.Debug("Failed to create publisher socket")
		return err
	}

	url := fmt.Sprintf("ipc://%s_point_publish", uid)

	err = p.socket.Listen(url)
	if err != nil {
		p.logger.Debug("Failed to listen on publisher socket")
		return err
	}

	return err
}

// Stop stops the Publisher.
func (p *Publisher) Stop() (err error) {
	if !p.started.Load() {
		p.logger.Debug("Publisher already stopped")
		return ErrAlreadyStopped
	}

	p.started.Store(false)

	err = p.socket.Close()
	if err != nil {
		p.logger.Debug("Failed to close publisher socket")
		return
	}

	return
}

// Publish publishes a message.
func (p *Publisher) Publish(message []byte) (err error) {
	if !p.started.Load() {
		p.logger.Debug("Publisher not started")
		return ErrNotStarted
	}

	err = p.socket.Send(message)
	if err != nil {
		p.logger.Debug("Failed to send message")
		return
	}

	return
}
