package loggerModel

type LogLevel int8
type LogLevelEnum map[int8]LogLevel

const (
	debugLevel = 1
	infoLevel  = 2
	warnLevel  = 3
	errorLevel = 4
	fatalLevel = 5
)

var Levels = LogLevelEnum{
	debugLevel: debugLevel,
	infoLevel:  infoLevel,
	warnLevel:  warnLevel,
	errorLevel: errorLevel,
	fatalLevel: fatalLevel,
}

func (e LogLevelEnum) Debug() LogLevel {
	return e[debugLevel]
}

func (e LogLevelEnum) Info() LogLevel {
	return e[infoLevel]
}

func (e LogLevelEnum) Warn() LogLevel {
	return e[warnLevel]
}

func (e LogLevelEnum) Error() LogLevel {
	return e[errorLevel]
}

func (e LogLevelEnum) Fatal() LogLevel {
	return e[fatalLevel]
}
