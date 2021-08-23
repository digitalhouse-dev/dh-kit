package response

import (
	"encoding/json"
	"net/http"

	"github.com/digitalhouse-dev/dh-kit/headers"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Errors  interface{}     `json:"errors,omitempty"`
	Headers headers.Headers `json:"-"`
}

const (
	badRequestMsg          = "Your request is in a bad format."
	forbiddenMsg           = "You are not authorized to perform the requested action."
	unauthorizedMsg        = "You are not authenticated to perform the requested action."
	notFoundMsg            = "The requested resource was not found."
	internalServerErrorMsg = "We encountered an error while processing your request."
	invalidInputMsg        = "There is some problem with the data you submitted."
)

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

// GetHeaders is
func (e ErrorResponse) GetHeaders() map[string]string {
	return e.Headers.Get()
}

// GetData return body for success and error interface for errors
func (e *ErrorResponse) GetData() interface{} {
	return e.Errors
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) Response {
	return makeErrorResponse(msg, internalServerErrorMsg, nil, http.StatusInternalServerError)
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string) Response {
	return makeErrorResponse(msg, notFoundMsg, nil, http.StatusNotFound)
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string) Response {
	return makeErrorResponse(msg, unauthorizedMsg, nil, http.StatusUnauthorized)
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg string) Response {
	return makeErrorResponse(msg, forbiddenMsg, nil, http.StatusForbidden)
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) Response {
	return makeErrorResponse(msg, badRequestMsg, nil, http.StatusBadRequest)
}

// InvalidInput creates a new error response representing a data validation error (HTTP 400) with error interface.
func InvalidInput(msg string, errors interface{}) Response {
	return makeErrorResponse(msg, invalidInputMsg, errors, http.StatusBadRequest)
}

func makeErrorResponse(msg string, defMsg string, errors interface{}, status int) Response {
	if msg == "" {
		msg = defMsg
	}
	headers := headers.New()
	return &ErrorResponse{
		Status:  status,
		Message: msg,
		Headers: *headers,
		Errors:  errors,
	}
}
