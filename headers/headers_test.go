package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorsHeadersByDefault(t *testing.T) {
	cors := make(map[string]string)
	cors["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"
	cors["Access-Control-Allow-Methods"] = "GET,OPTIONS,POST,PUT,PATCH"
	cors["Access-Control-Allow-Origin"] = "*"

	h := *New()
	assert.EqualValues(t, h.Get(), cors)
}

func TestSetHeaderAndCheckCors(t *testing.T) {
	h := *New()
	h.Add("content-type", "application/json charset=utf-8")
	assert.EqualValues(t, h.GetValueByKey("content-type"), "application/json charset=utf-8")
	assert.EqualValues(t, h.GetValueByKey("Access-Control-Allow-Origin"), "*")
}
