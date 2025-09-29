package test

import (
	"context"
	"dev-knowledge/infrastructure/logger"
	loggerModel "dev-knowledge/infrastructure/logger/model"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"time"
)

type LogServiceShould struct {
	suite.Suite
	stopCh        chan struct{}
	loggerService *LogServiceSpy
	logPublisher  *logger.LogPublisher
	ctx           context.Context
}

func (c *LogServiceShould) SetupTest() {
	c.stopCh = make(chan struct{})
	c.loggerService = GetLogServiceSpy(c.stopCh)
	c.logPublisher = logger.NewLogPublisher(c.loggerService.LogCh)
}

func (c *LogServiceShould) TearDownTest() {
	close(c.stopCh)
}

func (c *LogServiceShould) TestPublish_WrappedError_Success() {
	err := extErrors.New("test")

	c.logPublisher.LogError(c.ctx, err)
	time.Sleep(10 * time.Millisecond)

	received := c.loggerService.LastLogMsg
	assert.NotNil(c.T(), received)
	assert.Equal(c.T(), loggerModel.Levels.Error(), received.Level)
	assert.Contains(c.T(), received.Fields[0].String, "LogPublisher")

	assert.Equal(c.T(), 1, c.loggerService.NumLogReceived)
}

func (c *LogServiceShould) TestPublish_ManyErrors_Success() {
	err1 := errors.New("test1")
	err2 := errors.New("test2")

	c.logPublisher.LogError(c.ctx, err1, err2)
	time.Sleep(10 * time.Millisecond)

	assert.Equal(c.T(), 2, c.loggerService.NumLogReceived)
}
