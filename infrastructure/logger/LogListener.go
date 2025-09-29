package logger

import (
	logInterface "dev-knowledge/infrastructure/logger/interface"
	loggerModel "dev-knowledge/infrastructure/logger/model"
	"runtime"
)

const (
	inputChanSize = 100
)

type LogListener struct {
	logger    logInterface.Logger
	inputChan chan *loggerModel.LogData
	stopChan  chan struct{}
}

func NewLogListener(stopChan chan struct{}, logger logInterface.Logger) *LogListener {
	return &LogListener{
		inputChan: make(chan *loggerModel.LogData, inputChanSize),
		stopChan:  stopChan,
		logger:    logger,
	}
}

func (ls *LogListener) Start() {
	go ls.startWorker()
}

func (ls *LogListener) InputChan() chan *loggerModel.LogData {
	return ls.inputChan
}

func (ls *LogListener) startWorker() {
	for {
		select {
		case <-ls.stopChan:
			close(ls.inputChan)
			return
		case logData := <-ls.inputChan:
			ls.logger.LogMsg(logData)
		}
		runtime.Gosched()
	}
}
