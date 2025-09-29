package logInterface

import loggerModel "dev-knowledge/infrastructure/logger/model"

type Logger interface {
	LogMsg(data *loggerModel.LogData)
}
