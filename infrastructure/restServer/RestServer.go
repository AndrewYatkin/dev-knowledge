package restServer

import (
	"context"
	jwtServiceInterface "dev-knowledge/infrastructure/jwtService/interface"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	restServerInterface "dev-knowledge/infrastructure/restServer/interface"
	middleware "dev-knowledge/infrastructure/restServer/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

type ctxKey string

const RequestParamsKey ctxKey = "requestParams"

type FiberServer struct {
	server        *fiber.App
	logger        loggerInterface.Logger
	jwtMiddleware fiber.Handler
}

func NewFiberServer(logger loggerInterface.Logger, service jwtServiceInterface.JWTService) restServerInterface.Server {
	server := fiber.New(fiber.Config{
		ErrorHandler: middleware.NewErrorMiddleware(logger).Handler(),
	})
	server.Use(middleware.NewRequestMiddleware(logger).Handler())

	return &FiberServer{
		server:        server,
		logger:        logger,
		jwtMiddleware: middleware.NewJWTMiddleware(logger, service).Handler(),
	}
}

func (s *FiberServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {
	s.registerFiberRoute(method, path, httpHandlerFuncToFiberHandler(handler))
}

func (s *FiberServer) RegisterPrivateRoute(method, path string, handler http.HandlerFunc) {
	fiberHandler := httpHandlerFuncToFiberHandler(handler)
	s.registerFiberRoute(method, path, s.jwtMiddleware, fiberHandler)
}

func (s *FiberServer) registerFiberRoute(method, path string, handlers ...fiber.Handler) {
	switch method {
	case "GET":
		s.server.Get(path, handlers...)
	case "POST":
		s.server.Post(path, handlers...)
	case "PUT":
		s.server.Put(path, handlers...)
	case "DELETE":
		s.server.Delete(path, handlers...)
	default:
		panic("Unsupported method")
	}
}

func (s *FiberServer) registerFiberRouteWithMiddleware(method, path string, routeMiddleware fiber.Handler, handler fiber.Handler) {
	combineHandler := func(c *fiber.Ctx) error {
		if err := routeMiddleware(c); err != nil {
			return err
		}
		return handler(c)
	}
	s.registerFiberRoute(method, path, combineHandler)
}

func (s *FiberServer) Start(address string) error {
	return s.server.Listen(address)
}

func httpHandlerFuncToFiberHandler(handler http.HandlerFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(http.Request)
		err := fasthttpadaptor.ConvertRequest(c.Context(), req, true)
		if err != nil {
			return err
		}

		params := c.AllParams()
		ctx := context.WithValue(c.UserContext(), RequestParamsKey, params)
		req = req.WithContext(ctx)

		rw := &ResponseWriter{ctx: c}

		handler(rw, req)

		return nil
	}
}
