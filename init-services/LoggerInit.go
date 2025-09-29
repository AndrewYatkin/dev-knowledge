package initServices

import (
	"dev-knowledge/infrastructure/logger"
	logInterface "dev-knowledge/infrastructure/logger/interface"
)

type LoggerInit struct {
	LogPublisher logInterface.LogPublisher
	logListener  *logger.LogListener
}

func NewLoggerInit() *LoggerInit {
	return &LoggerInit{}
}

func (i *LoggerInit) Init(appID string, env string, stopChan chan struct{}) logInterface.LogPublisher {

	zapLogger := logger.NewZapLogger(appID, env)

	i.logListener = logger.NewLogListener(stopChan, zapLogger)
	i.logListener.Start()

	i.LogPublisher = logger.NewLogPublisher(i.logListener.InputChan())

	return i.LogPublisher
}
