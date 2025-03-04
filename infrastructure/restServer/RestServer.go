package restServer

import (
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	restServerInterface "dev-knowledge/infrastructure/restServer/interface"
	middleware "dev-knowledge/infrastructure/restServer/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

type FiberServer struct {
	server *fiber.App
	logger loggerInterface.Logger
}

func NewFiberServer(logger loggerInterface.Logger) restServerInterface.Server {
	server := fiber.New(fiber.Config{
		ErrorHandler: middleware.NewErrorMiddleware(logger).Handler(),
	})
	server.Use(middleware.NewRequestMiddleware(logger).Handler())

	return &FiberServer{
		server: server,
		logger: logger,
	}
}

func (s *FiberServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {
	s.registerFiberRoute(method, path, httpHandlerFuncToFiberHandler(handler))
}

func (s *FiberServer) registerFiberRoute(method, path string, handler fiber.Handler) {
	switch method {
	case "GET":
		s.server.Get(path, handler)
	case "POST":
		s.server.Post(path, handler)
	case "PUT":
		s.server.Put(path, handler)
	case "DELETE":
		s.server.Delete(path, handler)
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
		fasthttpadaptor.NewFastHTTPHandlerFunc(handler)(c.Context())
		return nil
	}
}
