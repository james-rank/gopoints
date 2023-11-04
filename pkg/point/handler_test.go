package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTest struct {
	suite.Suite
	assert *assert.Assertions
	sub    *Handler
}

func (t *HandlerTest) SetupTest() {
	t.assert = assert.New(t.T())
	t.sub = Instance()
	t.assert.NotNil(t.sub)
	t.assert.NotNil(t.sub.points)
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTest))
}

func (t *HandlerTest) TestAddPoint() {
	point := &mockPoint{}
	t.sub.AddPoint("test", point)
	t.assert.Equal(1, len(t.sub.points))
	t.assert.Equal(point, t.sub.points["test"])
}

func (t *HandlerTest) TestRemovePoint() {
	point := &mockPoint{}
	t.sub.AddPoint("test", point)
	t.assert.Equal(1, len(t.sub.points))
	t.sub.RemovePoint("test")
	t.assert.Equal(0, len(t.sub.points))
}

func (t *HandlerTest) TestExecutePoint() {
	point := &mockPoint{}
	point.On("Execute").Return()
	t.sub.AddPoint("test", point)
	t.sub.ExecutePoint("test")
	point.AssertCalled(t.T(), "Execute")
}

func (t *HandlerTest) TestResumePoint() {
	point := &mockPoint{}
	point.On("Resume").Return()
	t.sub.AddPoint("test", point)
	t.sub.ResumePoint("test")
	point.AssertCalled(t.T(), "Resume")
}
