package messaging

import (
	"fmt"

	"github.com/james-rank/points/internal/pkg/log"
	"go.nanomsg.org/mangos/v3/protocol/req"
	_ "go.nanomsg.org/mangos/v3/transport/all" // Required for transport registration.
)

var newReqSocket = req.NewSocket

// Requester is a struct that implements the Requester interface.
type Requester struct {
	logger logger
}

// NewRequester creates a new Requester.
func NewRequester() *Requester {
	return &Requester{
		logger: log.NewLogger(),
	}
}

// WithLogger sets the logger for the Requester.
func (r *Requester) WithLogger(logger logger) *Requester {
	r.logger = logger
	return r
}

// SendRequest sends a request to a replier.
func (r *Requester) SendRequest(buf []byte, uid string) (msg []byte, err error) {
	var sock socket

	if sock, err = newReqSocket(); err != nil {
		r.logger.Debug("Failed to create request socket")
		return nil, err
	}

	url := fmt.Sprintf("ipc://%s_point_replier", uid)

	if err = sock.Dial(url); err != nil {
		r.logger.Debug("Failed to dial on request socket")
		return nil, err
	}

	if err = sock.Send(buf); err != nil {
		r.logger.Debug("Failed to send request")
		return nil, err
	}

	if msg, err = sock.Recv(); err != nil {
		r.logger.Debug("Failed to receive reply")
		return nil, err
	}

	if err = sock.Close(); err != nil {
		r.logger.Debug("Failed to close request socket")
	}

	return msg, nil
}
