package test

import loggerModel "dev-knowledge/infrastructure/logger/model"

type LogServiceSpy struct {
	stopCh         chan struct{}
	LogCh          chan *loggerModel.LogData
	LastLogMsg     *loggerModel.LogData
	NumLogReceived int
}

func GetLogServiceSpy(stopCh chan struct{}) *LogServiceSpy {
	logger := &LogServiceSpy{
		stopCh: stopCh,
		LogCh:  make(chan *loggerModel.LogData, 5),
	}

	go logger.startWorker()

	return logger
}

func (s *LogServiceSpy) startWorker() {
	for {
		select {
		case <-s.stopCh:
			return
		case msg := <-s.LogCh:
			s.LastLogMsg = msg
			s.NumLogReceived++
		}
	}
}
