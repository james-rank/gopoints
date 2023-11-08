package handler

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
