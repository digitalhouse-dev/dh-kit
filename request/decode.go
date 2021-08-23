package request

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/digitalhouse-dev/dh-kit/response"
)

func DecodeMap(m map[string]string, s interface{}) error {
	for k, v := range m {
		err := setField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value string) error {
	fv := getFieldValueByTag(name, "json", obj)
	if !fv.IsValid() || !fv.CanSet() {
		return nil
	}

	ft := fv.Type()
	val := reflect.ValueOf(value)

	if fv.Kind() == reflect.Int {
		i, err := strconv.Atoi(value)
		if err != nil {
			return response.BadRequest("")
		}
		val = reflect.ValueOf(i)
		fv.Set(val)
		return nil
	}

	if fv.Kind() == reflect.Slice {
		s := strings.Split(value, ",")
		switch ft.Elem().Kind() {
		case reflect.String:
			val = reflect.ValueOf(s)
			fv.Set(val)
		case reflect.Int:
			var items []int
			for _, v := range s {
				i, err := strconv.Atoi(v)
				if err != nil {
					return response.BadRequest("")
				}
				items = append(items, i)
				val = reflect.ValueOf(items)
				fv.Set(val)
			}
		}
	}

	if ft != val.Type() {
		return nil
	}

	fv.Set(val)
	return nil
}

func getFieldValueByTag(tag, key string, s interface{}) reflect.Value {
	val := reflect.ValueOf(s).Elem()
	if val.Kind() != reflect.Struct {
		return reflect.Value{}
	}
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		v := strings.Split(typeField.Tag.Get(key), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		if v == tag {
			return val.Field(i)
		}
	}
	return reflect.Value{}
}
