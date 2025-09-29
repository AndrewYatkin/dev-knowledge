package restMiddleware

import (
	logInterface "dev-knowledge/infrastructure/logger/interface"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type RequestMiddleware struct {
	logger logInterface.LogPublisher
}

func NewRequestMiddleware(logger logInterface.LogPublisher) *RequestMiddleware {
	return &RequestMiddleware{
		logger: logger,
	}
}

func (r *RequestMiddleware) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		r.logger.LogInfo(c.Context(), fmt.Sprintf("Request: Method=%s, Path=%s",
			c.Method(), c.Path()))
		return c.Next()
	}
}
