package restMiddleware

import (
	logInterface "dev-knowledge/infrastructure/logger/interface"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ErrorMiddleware struct {
	logger logInterface.LogPublisher
}

func NewErrorMiddleware(logger logInterface.LogPublisher) *ErrorMiddleware {
	return &ErrorMiddleware{
		logger: logger,
	}
}

func (r *ErrorMiddleware) Handler() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		r.logger.LogError(c.Context(), fmt.Errorf("error: %v, Path: %s, Method: %s",
			err, c.Path(), c.Method()))
		return c.Status(code).JSON(fiber.Map{"error": err.Error()})
	}
}
