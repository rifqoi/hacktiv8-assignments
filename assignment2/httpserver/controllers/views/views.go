package views

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func BadRequestError(err error, msg string) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "BAD_REQUEST",
		Error:   msg,
	}
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err.Error(),
	}
}

func DataConflict(err error) *Response {
	return &Response{
		Status:  http.StatusConflict,
		Message: "DUPLICATE_DATA",
		Error:   err.Error(),
	}
}

func SuccessResponse(payload interface{}, msg string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: msg,
		Payload: payload,
	}
}

func FailedNotFound(err error) *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: "NOT_FOUND",
		Error:   err.Error(),
	}
}
