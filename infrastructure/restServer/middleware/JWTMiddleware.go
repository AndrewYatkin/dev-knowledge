package restMiddleware

import (
	jwtServiceInterface "dev-knowledge/infrastructure/jwtService/interface"
	logInterface "dev-knowledge/infrastructure/logger/interface"
	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware struct {
	logger     logInterface.LogPublisher
	jwtService jwtServiceInterface.JWTService
}

func NewJWTMiddleware(logger logInterface.LogPublisher, jwtService jwtServiceInterface.JWTService) *JWTMiddleware {
	return &JWTMiddleware{
		logger:     logger,
		jwtService: jwtService,
	}
}

func (r *JWTMiddleware) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := r.extractAuthToken(c)
		if err != nil {
			return err
		}

		if !r.jwtService.Verify(token) {
			return fiber.ErrUnauthorized
		}

		ctx, err := r.jwtService.FillCtxWithParams(c.UserContext(), token)
		if err != nil {
			r.logger.LogError(c.Context(), err)
			return fiber.ErrUnauthorized
		}

		c.SetUserContext(ctx)

		return c.Next()
	}

}

func (r *JWTMiddleware) extractAuthToken(c *fiber.Ctx) (string, error) {
	token := c.Get("Authorization")
	if token == "" {
		return "", fiber.ErrUnauthorized
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	return token, nil
}
