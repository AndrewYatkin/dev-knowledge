package restModel

type ErrorResponseData struct {
	HttpCode int    `json:"responseCode"`
	Text     string `json:"text"`
}

func NewErrorResponseData(httpCode int, text string) *ErrorResponseData {
	return &ErrorResponseData{
		HttpCode: httpCode,
		Text:     text,
	}
}
