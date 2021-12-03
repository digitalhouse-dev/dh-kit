package response

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse_Error(t *testing.T) {
	e := ErrorResponse{
		Message: "abc",
	}
	assert.Equal(t, "abc", e.Error())
}

func TestErrorResponse_StatusCode(t *testing.T) {
	e := ErrorResponse{
		Status: 400,
	}
	assert.Equal(t, 400, e.StatusCode())
}

func TestInternalServerError(t *testing.T) {
	res := InternalServerError("test")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode())
	assert.Equal(t, "test", res.Error())
	res = InternalServerError("")
	assert.NotEmpty(t, res.Error())
}

func TestNotFound(t *testing.T) {
	res := NotFound("test")
	assert.Equal(t, http.StatusNotFound, res.StatusCode())
	assert.Equal(t, "test", res.Error())
	res = NotFound("")
	assert.NotEmpty(t, res.Error())
}

func TestUnauthorized(t *testing.T) {
	res := Unauthorized("test")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode())
	assert.Equal(t, "test", res.Error())
	res = Unauthorized("")
	assert.NotEmpty(t, res.Error())
}

func TestForbidden(t *testing.T) {
	res := Forbidden("test")
	assert.Equal(t, http.StatusForbidden, res.StatusCode())
	assert.Equal(t, "test", res.Error())
	res = Forbidden("")
	assert.NotEmpty(t, res.Error())
}

func TestBadRequest(t *testing.T) {
	res := BadRequest("test")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	assert.Equal(t, "test", res.Error())
	res = BadRequest("")
	assert.NotEmpty(t, res.Error())
}

func TestResponseInvalidInput(t *testing.T) {

	type testStruct struct {
		value1 string
		value2 int
	}
	s := testStruct{"TestResponseInvalidInput unittest", 31}

	res := InvalidInput("unit  test", s)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode())
	assert.Equal(t, "unit  test", res.Error())

	data := res.GetData().(testStruct)
	assert.Equal(t, data.value1, "TestResponseInvalidInput unittest")
	assert.Equal(t, data.value2, 31)

	res = InvalidInput("", nil)
	assert.NotEmpty(t, res.Error())
	assert.Nil(t, res.GetData())
}
