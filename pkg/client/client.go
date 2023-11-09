// Package client
package client

import (
	"errors"
	"time"

	"github.com/james-rank/points/internal/pkg/log"
	"github.com/james-rank/points/pkg/messaging"
	"github.com/james-rank/points/pkg/protos"
	"google.golang.org/protobuf/proto"
)

var (
	ErrNack       = errors.New("received NACK")
	ErrTimeout    = errors.New("timed out waiting for notification")
	ErrWrongPoint = errors.New("received notification for wrong point")
)

// Client is a client for the points system.
type Client struct {
	logger     logger
	requester  requester
	subscriber subscriber
}

// NewClient returns a new Client.
func NewClient() *Client {
	return &Client{
		logger:     log.NewLogger(),
		requester:  messaging.NewRequester(),
		subscriber: messaging.NewSubscriber(),
	}
}

// WithLogger sets the logger for the client.
func (c *Client) WithLogger(l logger) *Client {
	c.logger = l
	return c
}

// WithRequester sets the requester for the client.
func (c *Client) WithRequester(r requester) *Client {
	c.requester = r
	return c
}

// WithSubscriber sets the subscriber for the client.
func (c *Client) WithSubscriber(s subscriber) *Client {
	c.subscriber = s
	return c
}

func (c *Client) Start(uid string) error {
	err := c.subscriber.Start(uid)
	if err != nil {
		c.logger.Debug("Failed to start subscriber")
	}

	return err
}

func (c *Client) Stop() error {
	return c.subscriber.Stop()
}

// SendRequest sends a request to a replier.
func (c *Client) SendRequest(buf []byte, uid string) error {
	buf, err := c.requester.SendRequest(buf, uid)
	if err != nil {
		c.logger.Debug("Failed to send request")
		return err
	}

	reply := &protos.Reply{}

	err = proto.Unmarshal(buf, reply)
	if err != nil {
		c.logger.Debug("Failed to unmarshal reply")
		return err
	}

	if reply.Response == protos.Reply_RESPONSE_NACK {
		c.logger.Debug("Received NACK")
		return ErrNack
	}

	return nil
}

func (c *Client) WaitForNotification(pointName string, timeout time.Duration) error {
	select {
	case msg := <-c.subscriber.Subscribe():
		notification := &protos.Notification{}

		err := proto.Unmarshal(msg, notification)
		if err != nil {
			c.logger.Debug("Failed to unmarshal notification")
			return err
		}

		if notification.Name == pointName {
			return nil
		}

		return ErrWrongPoint

	case <-time.After(timeout):
		c.logger.Debug("Timed out waiting for notification")
		return ErrTimeout
	}
}
