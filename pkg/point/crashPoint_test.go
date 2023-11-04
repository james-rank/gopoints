package point

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CrashPointTest struct {
	suite.Suite
	assert *assert.Assertions
	sub    *CrashPoint
	log    *mockLogger
}

func (t *CrashPointTest) SetupTest() {
	t.assert = assert.New(t.T())
	t.log = &mockLogger{}
	t.sub = NewCrashPoint("test").WithLogger(t.log)
	t.assert.NotNil(t.sub)
}

func TestCrashPoint(t *testing.T) {
	suite.Run(t, new(CrashPointTest))
}

func (t *CrashPointTest) TestExecute() {
	t.log.On("Debug", "Point 'test' Crashing...").Return()
	defer func() {
		r := recover()
		t.assert.NotNil(r)
		t.assert.Equal(fmt.Sprintf("Point 'test' Crashed!"), r)
	}()
	t.sub.Execute()
}

func (t *CrashPointTest) TestResume() {
	t.sub.Resume()
}
