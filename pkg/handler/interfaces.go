package handler

import (
	"context"
	"sync"

	"github.com/james-rank/points/pkg/messaging"
)

type logger interface {
	Debug(string)
}

type point interface {
	Execute()
}

type publisher interface {
	Start(string) error
	Stop() error
	Publish([]byte) error
}

type replier interface {
	Run(context.Context, *sync.WaitGroup, string, messaging.MsgFunc) error
}
