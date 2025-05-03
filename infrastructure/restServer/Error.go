package restServer

import "dev-knowledge/infrastructure/errors"

var (
	ErrURLParamsIsEmpty = errors.NewError("62106aaa-001", "URL Params is empty")
	ErrParseURLParams   = errors.NewError("62106aaa-002", "Can't parse url params")
)
