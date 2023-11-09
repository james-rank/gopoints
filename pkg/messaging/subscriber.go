package messaging

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/james-rank/points/internal/pkg/log"
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/sub"
	_ "go.nanomsg.org/mangos/v3/transport/all" // Required for transport registration.
)

const (
	subscriberQueueSize = 20
	waitBetweenRetries  = time.Millisecond * 100
)

var newSubSocket = sub.NewSocket

// Subscriber is a struct that implements the Subscriber interface.
type Subscriber struct {
	logger  logger
	socket  socket
	ctx     context.Context
	ch      chan []byte
	wg      *sync.WaitGroup
	started atomic.Bool
}

// NewSubscriber creates a new Subscriber.
func NewSubscriber() *Subscriber {
	return &Subscriber{
		logger:  log.NewLogger(),
		started: atomic.Bool{},
		ch:      make(chan []byte, subscriberQueueSize),
	}
}

// WithLogger sets the logger for the Subscriber.
func (s *Subscriber) WithLogger(logger logger) *Subscriber {
	s.logger = logger
	return s
}

func (s *Subscriber) messageLoop() {
	defer s.wg.Done()

	for s.started.Load() {
		select {
		case <-s.ctx.Done():
			s.logger.Debug("Subscriber context done")
			s.started.Store(false)

			err := s.socket.Close()
			if err != nil {
				s.logger.Debug("Failed to close subscriber socket")
				return
			}
		case <-time.After(waitBetweenRetries):
			msg, err := s.socket.Recv()
			if err != nil {
				s.logger.Debug("Failed to receive message from publisher socket: " + err.Error())
			} else {
				select {
				case s.ch <- msg:
				default:
					s.logger.Debug("Subscriber channel full")
				}
			}
		}
	}
}

// Start starts the Subscriber.
func (s *Subscriber) Start(uid string) (err error) {
	if s.started.Load() {
		s.logger.Debug("Subscriber already started")
		return ErrAlreadyStarted
	}

	s.socket, err = newSubSocket()
	if err != nil {
		s.logger.Debug("Failed to create subscriber socket")
		return err
	}

	url := fmt.Sprintf("ipc://%s_point_publish", uid)

	err = s.socket.Dial(url)
	if err != nil {
		s.logger.Debug("Failed to dial on subscriber socket")
		return err
	}

	// Empty byte array effectively subscribes to everything.
	err = s.socket.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		s.logger.Debug("Failed to set subscriber options")
		return err
	}

	s.started.Store(true)

	go s.messageLoop()

	return nil
}

// Stop stops the Subscriber.
func (s *Subscriber) Stop() error {
	if !s.started.Load() {
		s.logger.Debug("Subscriber already stopped")
		return ErrAlreadyStopped
	}

	s.started.Store(false)

	s.wg.Wait()

	return nil
}

// Subscribe returns the channel that the Subscriber publishes to.
func (s *Subscriber) Subscribe() <-chan []byte {
	return s.ch
}
