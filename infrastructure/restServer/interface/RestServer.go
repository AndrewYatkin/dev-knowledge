package restServerInterface

import "net/http"

type Server interface {
	RegisterPublicRoute(method, path string, handler http.HandlerFunc)
	Start(address string) error
}
