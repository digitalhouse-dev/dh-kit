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

// Success creates a new SuccessResponse instance.
// The message parameter refers to the message that will be use in the response body.
// The data parameter refers to the data key in the response body.
// The meta parameter refers to the meta key used in the response body.
// And the h parameter specifies the headers.
func Success(msg string, data interface{}, meta *meta.Meta, h *headers.Headers) Response {
	if h == nil {
		h = headers.New()
	}
	if msg == "" {
		msg = "Success request."
	}
	return &SuccessResponse{
		Message: msg,
		Status:  http.StatusOK,
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
