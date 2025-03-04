package restMiddleware

import (
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type RequestMiddleware struct {
	logger loggerInterface.Logger
}

func NewRequestMiddleware(logger loggerInterface.Logger) *RequestMiddleware {
	return &RequestMiddleware{
		logger: logger,
	}
}

func (r *RequestMiddleware) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		r.logger.Info(c.Context(), fmt.Sprintf("Request: Method=%s, Path=%s",
			c.Method(), c.Path()))
		return c.Next()
	}
}
