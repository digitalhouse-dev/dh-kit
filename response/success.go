package response

import (
	"encoding/json"
	"net/http"

	"github.com/digitalhouse-dev/dh-kit/headers"
	"github.com/digitalhouse-dev/dh-kit/meta"
)

type SuccessResponse struct {
	Message string           `json:"message"`
	Status  int              `json:"status"`
	Data    interface{}      `json:"data"`
	Meta    *meta.Meta       `json:"meta,omitempty"`
	Headers *headers.Headers `json:"-"`
}

// OK creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func OK(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if msg == "" {
		msg = "Success request."
	}

	return success(msg, data, meta, h, http.StatusOK)
}

// Created creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func Created(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if msg == "" {
		msg = "Created request."
	}

	return success(msg, data, meta, h, http.StatusCreated)
}

// Accepted creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func Accepted(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if msg == "" {
		msg = "Accepted request."
	}

	return success(msg, data, meta, h, http.StatusAccepted)
}

// NonAuthoritativeInfo creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func NonAuthoritativeInfo(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if msg == "" {
		msg = "Non-Authoritative Information request."
	}

	return success(msg, data, meta, h, http.StatusNonAuthoritativeInfo)
}

// NoContent creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func NoContent(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if msg == "" {
		msg = "No Content request."
	}

	return success(msg, data, meta, h, http.StatusNoContent)
}

// ResetContent creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func ResetContent(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if msg == "" {
		msg = "Reset Content request."
	}

	return success(msg, data, meta, h, http.StatusResetContent)
}

// PartialContent creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func PartialContent(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if msg == "" {
		msg = "Partial Content request."
	}

	return success(msg, data, meta, h, http.StatusPartialContent)
}

func success(msg string, data interface{}, meta *meta.Meta, h *headers.Headers, code int) Response {
	if h == nil {
		h = headers.New()
	}

	return &SuccessResponse{
		Message: msg,
		Status:  code,
		Data:    data,
		Meta:    meta,
		Headers: h,
	}
}

func (s *SuccessResponse) Error() string {
	return ""
}

func (s *SuccessResponse) StatusCode() int {
	return s.Status
}

func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SuccessResponse) GetHeaders() map[string]string {
	return s.Headers.Get()
}

// GetData return body for success and error interface for errors
func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}
