package logger

import (
	loggerModel "dev-knowledge/infrastructure/logger/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	hostTag = "service_name"
	envTag  = "env"
	timeTag = "timestamp"
)

type ZapLogger struct {
	*zap.Logger
}

func NewZapLogger(appID, env string) *ZapLogger {
	config := getEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	zapLogger := zap.New(core, createOptions(appID, env))

	return &ZapLogger{
		zapLogger,
	}
}

func (l *ZapLogger) LogMsg(logData *loggerModel.LogData) {
	resFields := l.getPayloadFields(logData)

	switch logData.Level {
	case loggerModel.Levels.Error():
		l.Error(logData.Msg, resFields...)
	case loggerModel.Levels.Warn():
		l.Warn(logData.Msg, resFields...)
	case loggerModel.Levels.Info():
		l.Info(logData.Msg, resFields...)
	case loggerModel.Levels.Debug():
		l.Debug(logData.Msg, resFields...)
	case loggerModel.Levels.Fatal():
		l.Fatal(logData.Msg, resFields...)
	}
}

func (l *ZapLogger) getPayloadFields(logData *loggerModel.LogData) []zap.Field {
	var resFields []zap.Field
	for _, f := range logData.Fields {
		if f.Integer != 0 || f.Key == "integer_zero_value" {
			resFields = append(resFields, zap.Int(f.Key, f.Integer))
		}
		if f.Float != 0.0 || f.Key == "float_zero_value" {
			resFields = append(resFields, zap.Float64(f.Key, f.Float))
		}
		if f.String != "" {
			resFields = append(resFields, zap.String(f.Key, f.String))
		}
	}
	return resFields
}

func createOptions(appID, env string) zap.Option {
	return zap.Fields(
		zap.String(hostTag, appID),
		zap.String(envTag, env),
	)
}

func getEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = timeTag
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	return config
}
