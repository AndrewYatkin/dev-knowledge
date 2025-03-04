package response

import (
	"bytes"
	"context"
	"dev-knowledge/infrastructure/errors"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	restServerInterface "dev-knowledge/infrastructure/restServer/interface"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponseService struct {
	errorResolver restServerInterface.ErrorResolver
	logger        loggerInterface.Logger
}

func NewErrorResponseService(
	errorResolver restServerInterface.ErrorResolver,
	logger loggerInterface.Logger,
) (*ErrorResponseService, error) {
	if errorResolver == nil {
		return nil, errors.NewError("SYS", "ErrorResolver is required")
	}
	if logger == nil {
		return nil, errors.NewError("SYS", "Logger is required")
	}

	return &ErrorResponseService{
		logger:        logger,
		errorResolver: errorResolver,
	}, nil
}

func (s *ErrorResponseService) ErrorsResponse(w http.ResponseWriter, r *http.Request, errors []error) {
	errorResponse := s.createErrorResponse(errors...)
	s.writeErrorResponse(r.Context(), w, errorResponse, errorResponse.FirstHttpCode())
}

func (s *ErrorResponseService) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse := s.createErrorResponse(err)
	s.writeErrorResponse(r.Context(), w, errorResponse, errorResponse.FirstHttpCode())
}

func (s *ErrorResponseService) createErrorResponse(errs ...error) ErrorResponse {
	errorsData := make([]ErrorResponseData, 0, len(errs))

	for _, err := range errs {
		switch err := err.(type) {
		case *errors.Errors:
			for _, errItem := range err.ToArray() {
				errData := s.createErrorResponseData(errItem)
				errorsData = append(errorsData, errData)
			}
		default:
			errData := s.createErrorResponseData(err)
			errorsData = append(errorsData, errData)
		}
	}

	return NewErrorResponse(errorsData)
}

func (s *ErrorResponseService) writeErrorResponse(ctx context.Context, w http.ResponseWriter, resp ErrorResponse, code int) {
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logError(ctx, "JSON marshal failed", err)
		return
	}

	w.Header().Set(HeaderContentType, "application/json; charset=utf-8")
	w.Header().Set(HeaderXContentTypeOptions, "nosniff")
	w.WriteHeader(code)

	prettyJSON, err := s.prettyJSON(body)
	if err != nil {
		s.logError(ctx, "err prettify response", err)
	}
	_, err = w.Write(prettyJSON)
	if err != nil {
		s.logError(ctx, "Response Writer Error", err)
		return
	}
}

func (s *ErrorResponseService) createErrorResponseData(err error) ErrorResponseData {
	responseCode := s.errorResolver.GetHttpCode(err)
	errorCode := s.errorResolver.GetErrorCode(err)
	errorText := s.errorResolver.GetErrorText(err)

	return ErrorResponseData{
		HttpCode:  responseCode,
		ErrorCode: errorCode,
		Text:      errorText,
	}
}

func (s *ErrorResponseService) prettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func (s *ErrorResponseService) logError(ctx context.Context, msg string, err error) {
	s.logger.Error(ctx, fmt.Errorf("%s: %v", msg, err))
}
