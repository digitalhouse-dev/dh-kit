package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetField(t *testing.T) {
	t.Run("successParsingInt", func(t *testing.T) {
		type Test struct {
			Number int `json:"number"`
		}
		obj := Test{}
		want := Test{
			Number: 1,
		}
		mapping := make(map[string]string)
		mapping["number"] = "1"
		err := DecodeMap(mapping, &obj)
		assert.Nil(t, err)
		assert.Equal(t, want, obj)
	})
	t.Run("successParsingBool", func(t *testing.T) {
		type Test struct {
			Boolean bool `json:"number"`
		}
		obj := Test{}
		want := Test{
			Boolean: true,
		}
		mapping := make(map[string]string)
		mapping["number"] = "true"
		err := DecodeMap(mapping, &obj)
		assert.Nil(t, err)
		assert.Equal(t, want, obj)
	})
	t.Run("successParsingBoolPointer", func(t *testing.T) {
		type Test struct {
			Boolean *bool `json:"boolean"`
		}
		obj := Test{}
		value := true
		want := Test{
			Boolean: &value,
		}
		mapping := make(map[string]string)
		mapping["boolean"] = "true"
		err := DecodeMap(mapping, &obj)
		assert.Nil(t, err)
		assert.Equal(t, *want.Boolean, *obj.Boolean)
	})
	t.Run("successParsingNilBoolPointer", func(t *testing.T) {
		type Test struct {
			Boolean *bool `json:"boolean"`
		}
		got := Test{}
		mapping := make(map[string]string)
		err := DecodeMap(mapping, &got)
		assert.Nil(t, err)
		assert.Nil(t, got.Boolean)
	})
}
