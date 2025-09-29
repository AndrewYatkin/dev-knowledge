package logger

import (
	"context"
	commonErrors "dev-knowledge/infrastructure/errors"
	loggerModel "dev-knowledge/infrastructure/logger/model"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

const maxStackTraceLevel = 21

var errorForLogIsNil = errors.New("unexpected behaviour in logPublisher: was received nil-error for logging")

type LogPublisher struct {
	logChan chan *loggerModel.LogData
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogPublisher(logChan chan *loggerModel.LogData) *LogPublisher {
	return &LogPublisher{logChan: logChan}
}

func (l *LogPublisher) LogDebug(ctx context.Context, messages ...string) {
	for _, message := range messages {
		logData := loggerModel.DebugLogData(ctx, message)
		go l.sendLogData(logData)
	}
}

func (l *LogPublisher) LogInfo(ctx context.Context, messages ...string) {
	for _, message := range messages {
		logData := loggerModel.InfoLogData(ctx, message)
		go l.sendLogData(logData)
	}
}

func (l *LogPublisher) LogWarn(ctx context.Context, messages ...string) {
	for _, message := range messages {
		logData := loggerModel.WarnLogData(ctx, message)
		go l.sendLogData(logData)
	}
}

func (l *LogPublisher) LogError(ctx context.Context, errs ...error) {
	for _, err := range errs {
		l.processLogError(ctx, err)
	}
}

func (l *LogPublisher) processLogError(ctx context.Context, err error) {
	logData := createErrLogData(ctx, err)
	go l.sendLogData(logData)
}

func getFileNames(errWithStack error) []string {
	stackTraceErr := errWithStack.(stackTracer)
	stackTrace := stackTraceErr.StackTrace()

	var fileNames []string
	if len(stackTrace) > 0 {
		for i := 1; i < len(stackTrace) && i < maxStackTraceLevel; i++ {
			fileNames = append(fileNames, fmt.Sprintf("%s:%d", stackTrace[i], stackTrace[i]))
		}
	}

	return fileNames
}

func createErrLogData(ctx context.Context, err error) *loggerModel.LogData {
	if err == nil {
		return &loggerModel.LogData{
			Ctx:    ctx,
			Msg:    errorForLogIsNil.Error(),
			Fields: []*loggerModel.LogField{},
			Level:  loggerModel.Levels.Error(),
		}
	}

	logLevel := getLogLevelFromError(err)
	errWithStack := errors.WithStack(err)
	fileNames := getFileNames(errWithStack)

	logFields := []*loggerModel.LogField{
		{
			Key:    loggerModel.FieldFileNameKey,
			String: strings.Join(fileNames, " <- "),
		},
	}

	return &loggerModel.LogData{
		Ctx:    ctx,
		Msg:    err.Error(),
		Fields: logFields,
		Level:  logLevel,
	}
}

func getLogLevelFromError(err error) loggerModel.LogLevel {
	switch err.(type) {
	case *commonErrors.Error:
		return getLogLevelFromErrorLevel(err.(*commonErrors.Error).Level())
	default:
		return loggerModel.Levels.Error()
	}
}

func getLogLevelFromErrorLevel(errorLevel commonErrors.ErrorLevel) loggerModel.LogLevel {
	switch errorLevel {
	case commonErrors.Levels.Info():
		return loggerModel.Levels.Info()
	case commonErrors.Levels.Warn():
		return loggerModel.Levels.Warn()
	case commonErrors.Levels.Error(),
		commonErrors.Levels.Critical():
		return loggerModel.Levels.Error()
	default:
		return loggerModel.Levels.Error()

	}
}

func (l *LogPublisher) sendLogData(logData *loggerModel.LogData) {
	l.logChan <- logData
}
