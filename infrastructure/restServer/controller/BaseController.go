package restServerController

import (
	"dev-knowledge/infrastructure/errors"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	restServerInterface "dev-knowledge/infrastructure/restServer/interface"
	"dev-knowledge/infrastructure/restServer/response"
	"io"
	"net/http"
)

type BaseController struct {
	responseService *response.ResponseService
	logPublisher    loggerInterface.Logger
}

func NewBaseController(
	responseService *response.ResponseService,
	logger loggerInterface.Logger,
) (*BaseController, error) {
	if responseService == nil {
		return nil, errors.NewError("SYS", "ResponseService is required")
	}
	if logger == nil {
		return nil, errors.NewError("SYS", "Logger is required")
	}

	return &BaseController{
		responseService: responseService,
		logPublisher:    logger,
	}, nil
}

func (bc *BaseController) FillReqModel(r *http.Request, reqModel restServerInterface.RequestModel) error {
	requestBody, err := bc.GetReqBody(r)
	if err != nil {
		return err
	}

	err = reqModel.FillFromBytes(requestBody)
	if err != nil {
		return response.ErrUnmarshalRequest(err.Error())
	}

	return err
}

func (bc *BaseController) GetReqBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, response.ErrUnmarshalRequest("Request body is nil")
	}
	return io.ReadAll(r.Body)
}

func (bc *BaseController) Response(w http.ResponseWriter, r *http.Request, result []byte, responseCode int) {
	bc.responseService.Response(w, r, result, responseCode)
}

func (bc *BaseController) JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}, responseCode int) {
	bc.responseService.JSONResponse(w, r, result, responseCode)
}

func (bc *BaseController) StreamResponse(w http.ResponseWriter, r *http.Request, response *http.Response) {
	bc.responseService.StreamResponse(w, r, response)
}

func (bc *BaseController) ErrorsResponse(w http.ResponseWriter, r *http.Request, errors []error) {
	bc.responseService.ErrorsResponse(w, r, errors)
}

func (bc *BaseController) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	bc.responseService.ErrorResponse(w, r, err)
}
