package messaging

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/james-rank/points/internal/pkg/log"
	"go.nanomsg.org/mangos/v3/protocol/rep"
	_ "go.nanomsg.org/mangos/v3/transport/all" // Required for transport registration.
)

var newRepSocket = rep.NewSocket

// MsgFunc is a function that processes a message.
type MsgFunc func([]byte) []byte

// Replier is a struct that implements the Replier interface.
type Replier struct {
	logger  logger
	sock    socket
	running atomic.Bool
}

// NewReplier creates a new Replier.
func NewReplier() *Replier {
	return &Replier{
		logger: log.NewLogger(),
	}
}

// WithLogger sets the logger for the Replier.
func (r *Replier) WithLogger(logger logger) *Replier {
	r.logger = logger
	return r
}

func (r *Replier) messageLoop(processMsg MsgFunc) {
	for r.running.Load() {
		msg, err := r.sock.Recv()
		if err != nil {
			r.logger.Debug("Failed to receive request")
		} else {
			reply := processMsg(msg)
			err = r.sock.Send(reply)
			if err != nil {
				r.logger.Debug("Failed to send reply")
			}
		}
	}
}

// Run starts the Replier.
func (r *Replier) Run(ctx context.Context, wg *sync.WaitGroup, uid string, processMsg MsgFunc) (err error) {
	defer wg.Done()

	if r.sock, err = newRepSocket(); err != nil {
		r.logger.Debug("Failed to create reply socket")
		return err
	}

	url := fmt.Sprintf("ipc://%s_point_replier", uid)

	if err = r.sock.Listen(url); err != nil {
		r.logger.Debug("Failed to listen on reply socket")
		return err
	}

	r.running.Store(true)

	go r.messageLoop(processMsg)

	<-ctx.Done()
	r.logger.Debug("Replier context done")

	r.running.Store(false)

	err = r.sock.Close()
	if err != nil {
		r.logger.Debug("Failed to close reply socket")
	}

	return nil
}
