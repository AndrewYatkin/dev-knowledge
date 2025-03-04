package response

import (
	"context"
	"dev-knowledge/infrastructure/errors"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	HeaderContentType         = "Content-Type"
	HeaderContentLength       = "Content-Length"
	HeaderXContentTypeOptions = "X-Content-Type-Options"
)

type ResponseService struct {
	errorResponseService *ErrorResponseService
	logPublisher         loggerInterface.Logger
}

func NewResponseService(
	errorResponseService *ErrorResponseService,
	logger loggerInterface.Logger,
) (*ResponseService, error) {
	if logger == nil {
		return nil, errors.NewError("SYS", "Logger is required")
	}
	if errorResponseService == nil {
		return nil, errors.NewError("SYS", "ErrorResponseService is required")
	}

	return &ResponseService{
		errorResponseService: errorResponseService,
		logPublisher:         logger,
	}, nil
}

func (s *ResponseService) JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}, responseCode int) {
	body, err := s.marshalBody(result)
	if err != nil {
		s.logError(r.Context(), "Marshal JSON error", err)
		s.ErrorResponse(w, r, ErrWriteResponse)
		return
	}
	s.Response(w, r, body, responseCode)
}

func (s *ResponseService) Response(w http.ResponseWriter, r *http.Request, result []byte, responseCode int) {
	w.Header().Set(HeaderContentType, "application/json; charset=utf-8")
	w.Header().Set(HeaderXContentTypeOptions, "nosniff")
	w.WriteHeader(responseCode)
	if responseCode != http.StatusNoContent {
		_, err := w.Write(result)
		if err != nil {
			s.logError(r.Context(), "Response Writer Error", err)
			return
		}
	}
}

func (s *ResponseService) StreamResponse(w http.ResponseWriter, r *http.Request, resp *http.Response) {
	defer resp.Body.Close()
	w.Header().Set(HeaderContentType, r.Header.Get(HeaderContentType))
	w.Header().Set(HeaderContentLength, r.Header.Get(HeaderContentLength))

	_, err := io.Copy(w, resp.Body)
	if err != nil {
		s.logError(r.Context(), "io.Copy for body failed", err)
	}
}

func (s *ResponseService) ErrorsResponse(w http.ResponseWriter, r *http.Request, errors []error) {
	s.errorResponseService.ErrorsResponse(w, r, errors)
}

func (s *ResponseService) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	s.errorResponseService.ErrorResponse(w, r, err)
}

func (s *ResponseService) marshalBody(result interface{}) ([]byte, error) {
	if result == nil || result == "" {
		return []byte{}, nil
	}

	body, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *ResponseService) logError(ctx context.Context, msg string, err error) {
	s.logPublisher.Error(ctx, fmt.Errorf("%s: %v", msg, err))
}
