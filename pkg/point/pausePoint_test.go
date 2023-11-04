package point

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PausePointTest struct {
	suite.Suite
	assert *assert.Assertions
	sub    *PausePoint
	log    *mockLogger
}

func (t *PausePointTest) SetupTest() {
	t.assert = assert.New(t.T())
	t.log = &mockLogger{}
	t.sub = NewPausePoint("test").WithLogger(t.log)
	t.assert.NotNil(t.sub)
}

func TestPausePoint(t *testing.T) {
	suite.Run(t, new(PausePointTest))
}

func (t *PausePointTest) TestExecuteSuccess() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	t.log.On("Debug", "Point 'test' Pausing...").Return()
	t.log.On("Debug", "Point 'test' Resuming...").Return()

	go func() {
		defer wg.Done()
		t.sub.Execute()
	}()

	t.sub.Resume()
	wg.Wait()
}

func (t *PausePointTest) TestExecuteTimeout() {
	t.sub = t.sub.WithTimeout(time.Nanosecond)
	t.log.On("Debug", "Point 'test' Pausing...").Return()
	t.log.On("Debug", "Timeout reached, Point 'test' resuming...").Return()

	t.sub.Execute()
}
