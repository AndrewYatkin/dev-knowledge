package logInterface

import "context"

type LogPublisher interface {
	LogError(ctx context.Context, err ...error)
	LogWarn(ctx context.Context, messages ...string)
	LogInfo(ctx context.Context, messages ...string)
	LogDebug(ctx context.Context, messages ...string)
}
