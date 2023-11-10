package todart

import (
	"encoding/json"
	"fmt"
	"github.com/any-call/gobase/util/myos"
	"github.com/any-call/gobase/util/mystr"
	"os"
	"reflect"
	"strings"
)

type Field struct {
	Name  string
	Type  string
	Value any
}

func jsonToModel(jsonStr string) (ret []Field, err error) {
	var jsonModel any
	if err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&jsonModel); err != nil {
		return nil, err
	}

	mapModel, ok := jsonModel.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("入参必须是一个map")
	}

	ret = make([]Field, 0, len(mapModel))
	for k, v := range mapModel {
		field := Field{Name: mystr.ToProperty(k), Value: v}
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.Type = "int?"
			break

		case reflect.Float32, reflect.Float64:
			field.Type = "double?"
			break

		case reflect.String:
			field.Type = "String?"
			break

		case reflect.Bool:
			field.Type = "bool?"
			break

		case reflect.Map:
			field.Type = "Map?"
			break

		case reflect.Slice, reflect.Array:
			field.Type = "list?"
			break

		default:
			field.Type = "Object?"
			break
		}

		ret = append(ret, field)
	}

	return ret, nil
}

func GenModel(destFile string, jsonStr string) error {
	listField, err := jsonToModel(jsonStr)
	if err != nil {
		return err
	}

	fileNameSlice := strings.Split(myos.Filename(destFile), ".")
	ClassName := fileNameSlice[0]
	dartModel := NewTempModel(mystr.ToTitle(ClassName), listField)
	return os.WriteFile(destFile, []byte(dartModel.ToCode()), 0777)
}
