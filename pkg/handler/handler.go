// Package handler is a package that defines the handler that runs in the program you want to test.
package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/james-rank/points/internal/pkg/log"
	"github.com/james-rank/points/pkg/messaging"
	pt "github.com/james-rank/points/pkg/point"
	"github.com/james-rank/points/pkg/protos"
	"google.golang.org/protobuf/proto"
)

var instance *Handler

func init() {
	instance = &Handler{
		logger:    log.NewLogger(),
		points:    make(map[string]point),
		publisher: messaging.NewPublisher(),
		replier:   messaging.NewReplier(),
	}
}

// Handler is a struct that holds the points.
type Handler struct {
	logger    logger
	points    map[string]point
	publisher publisher
	replier   replier
	ctx       context.Context
	cancel    context.CancelFunc
	wg        *sync.WaitGroup
}

// Instance returns the instance of the handler.
func Instance() *Handler {
	return instance
}

// WithLogger sets the logger for the handler.
func (h *Handler) WithLogger(logger logger) *Handler {
	h.logger = logger
	return h
}

// WithPublisher sets the publisher for the handler.
func (h *Handler) WithPublisher(publisher publisher) *Handler {
	h.publisher = publisher
	return h
}

// WithReplier sets the replier for the handler.
func (h *Handler) WithReplier(replier replier) *Handler {
	h.replier = replier
	return h
}

func (h *Handler) publishNotification(name string) {
	notification := &protos.Notification{Name: name}

	buf, err := proto.Marshal(notification)
	if err != nil {
		h.logger.Debug(fmt.Sprintf("failed to marshal notification for '%v': %v", name, err))
	} else {
		err = h.publisher.Publish(buf)
		if err != nil {
			h.logger.Debug(fmt.Sprintf("failed to publish notification for '%v': %v", name, err))
		}
	}
}

// ExecutePoint executes the point with the given name.
func (h *Handler) ExecutePoint(name string) {
	if p, ok := h.points[name]; ok {
		h.publishNotification(name)
		p.Execute()
	}
}

func (h *Handler) processAdd(request *protos.Request) []byte {
	if _, ok := h.points[request.Name]; ok {
		h.logger.Debug(fmt.Sprintf("point %s already exists", request.Name))
		return h.returnResponse(protos.Reply_RESPONSE_NACK)
	}

	switch request.Type {
	case protos.Request_POINT_PAUSE:
		h.points[request.Name] = pt.NewPausePoint(request.Name)
	case protos.Request_POINT_CRASH:
		h.points[request.Name] = pt.NewCrashPoint(request.Name)
	}

	return h.returnResponse(protos.Reply_RESPONSE_ACK)
}

func (h *Handler) processRemove(request *protos.Request) []byte {
	if _, ok := h.points[request.Name]; ok {
		delete(h.points, request.Name)
		return h.returnResponse(protos.Reply_RESPONSE_ACK)
	}

	h.logger.Debug(fmt.Sprintf("point %s does not exist", request.Name))

	return h.returnResponse(protos.Reply_RESPONSE_NACK)
}

func (h *Handler) processResume(request *protos.Request) []byte {
	if p, ok := h.points[request.Name]; ok {
		p.Execute()

		return h.returnResponse(protos.Reply_RESPONSE_ACK)
	}

	h.logger.Debug(fmt.Sprintf("point %s does not exist", request.Name))

	return h.returnResponse(protos.Reply_RESPONSE_NACK)
}

func (h *Handler) processRequest(request *protos.Request) []byte {
	switch request.Action {
	case protos.Request_ACTION_ADD:
		return h.processAdd(request)
	case protos.Request_ACTION_REMOVE:
		return h.processRemove(request)
	case protos.Request_ACTION_RESUME:
		return h.processResume(request)
	default:
		return h.returnResponse(protos.Reply_RESPONSE_NACK)
	}
}

func (h *Handler) returnResponse(resp protos.Reply_Response) []byte {
	reply := &protos.Reply{
		Response: resp,
	}

	msg, err := proto.Marshal(reply)
	if err != nil {
		h.logger.Debug(fmt.Sprintf("failed to marshal reply: %v", err))
	}

	return msg
}

func (h *Handler) handleRequests(buf []byte) []byte {
	var request protos.Request

	err := proto.Unmarshal(buf, &request)
	if err != nil {
		h.logger.Debug(fmt.Sprintf("failed to unmarshal request: %v", err))
		h.returnResponse(protos.Reply_RESPONSE_NACK)
	}

	return h.processRequest(&request)
}

// Start starts the handler.
func (h *Handler) Start(uid string) (err error) {
	err = h.publisher.Start(uid)
	if err != nil {
		return err
	}

	h.ctx, h.cancel = context.WithCancel(context.Background())
	h.wg = &sync.WaitGroup{}

	runFunc := func() {
		h.wg.Add(1)

		err := h.replier.Run(h.ctx, h.wg, uid, h.handleRequests)
		if err != nil {
			h.logger.Debug(fmt.Sprintf("failed to run replier: %v", err))
		}
	}

	go runFunc()

	return nil
}

// Stop stops the handler.
func (h *Handler) Stop() (err error) {
	err = h.publisher.Stop()
	if err != nil {
		return err
	}

	h.cancel()
	h.wg.Wait()

	return nil
}
