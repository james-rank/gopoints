package point

import (
	"fmt"
	"time"

	"github.com/james-rank/points/internal/pkg/log"
)

const timeout = 5 * time.Minute

// PausePoint is a point that pauses the program.
type PausePoint struct {
	ch      chan bool
	logger  logger
	name    string
	timeout time.Duration
}

// NewPausePoint returns a new PausePoint.
func NewPausePoint(name string) *PausePoint {
	return &PausePoint{
		ch:      make(chan bool),
		logger:  log.NewLogger(),
		name:    name,
		timeout: timeout,
	}
}

// WithLogger sets the logger for the point.
func (p *PausePoint) WithLogger(l logger) *PausePoint {
	p.logger = l
	return p
}

// WithTimeout sets the timeout for the point.
func (p *PausePoint) WithTimeout(t time.Duration) *PausePoint {
	p.timeout = t
	return p
}

// Execute executes the point.
func (p *PausePoint) Execute() {
	p.logger.Debug(fmt.Sprintf("Point '%s' Pausing...", p.name))

	select {
	case <-p.ch:
		p.logger.Debug(fmt.Sprintf("Point '%s' Resuming...", p.name))
	case <-time.After(p.timeout):
		p.logger.Debug(fmt.Sprintf("Timeout reached, Point '%s' resuming...", p.name))
	}
}

// Resume resumes the point.
func (p *PausePoint) Resume() {
	p.ch <- true
}
