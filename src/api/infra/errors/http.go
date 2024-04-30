package errors

import (
	"fmt"

	logging "website-conector/infra/logger"
)

type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewHttpError(code, message string) HttpError {
	logging.Error(code, message)
	return HttpError{Code: code, Message: message}
}

func (h HttpError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", h.Code, h.Message)
}
