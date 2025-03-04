package errors

func ToInfo(err error) *Error {
	return createErrorWithLevelFrom(err, Levels.Info())
}

func ToWarn(err error) *Error {
	return createErrorWithLevelFrom(err, Levels.Warn())
}

func ToError(err error) *Error {
	return createErrorWithLevelFrom(err, Levels.Error())
}

func ToCritical(err error) *Error {
	return createErrorWithLevelFrom(err, Levels.Critical())
}

func createErrorWithLevelFrom(err error, logLevel ErrorLevel) *Error {
	switch err := err.(type) {
	case *Error:
		return NewErrorWithLevel(err.Code(), err.Message(), logLevel)
	default:
		return NewErrorWithLevel(UnknownErrorCode, err.Error(), logLevel)
	}
}
