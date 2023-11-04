package point

import (
	"fmt"

	"github.com/james-rank/points/internal/pkg/log"
)

// CrashPoint is a point that crashes the program.
type CrashPoint struct {
	logger logger
	name   string
}

// NewCrashPoint returns a new CrashPoint.
func NewCrashPoint(name string) *CrashPoint {
	return &CrashPoint{
		logger: log.NewLogger(),
		name:   name,
	}
}

// WithLogger sets the logger for the point.
func (p *CrashPoint) WithLogger(l logger) *CrashPoint {
	p.logger = l
	return p
}

// Execute executes the point.
func (p *CrashPoint) Execute() {
	p.logger.Debug(fmt.Sprintf("Point '%s' Crashing...", p.name))
	panic(fmt.Sprintf("Point '%s' Crashed!", p.name))
}

// Resume resumes the point.
func (p *CrashPoint) Resume() {
	// Do nothing.
}
