package restServer

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ResponseWriter struct {
	c           *fiber.Ctx
	header      http.Header
	statusCode  int
	wroteHeader bool
}

func (w *ResponseWriter) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}
	return w.header
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	if !w.wroteHeader {
		w.writeHeader()
	}
	return w.c.Write(data)
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.writeHeader()
}

func (w *ResponseWriter) writeHeader() {
	if w.statusCode != 0 {
		w.c.Status(w.statusCode)
	} else {
		w.c.Status(http.StatusOK)
	}

	for k, vv := range w.header {
		for _, v := range vv {
			w.c.Response().Header.Add(k, v)
		}
	}
	w.wroteHeader = true
}
