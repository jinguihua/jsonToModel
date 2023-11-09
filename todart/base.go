package todart

import (
	"encoding/json"
	"fmt"
	"github.com/any-call/gobase/util/myos"
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
		field := Field{Name: k, Value: v}
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
			field.Type = "map?"
			break

		case reflect.Slice, reflect.Array:
			field.Type = "list?"
			break

		default:
			field.Type = "Object?"
			break
		}
	}

	return ret, nil
}

func GenModel(destFile string, jsonStr string) error {
	listField, err := jsonToModel(jsonStr)
	if err != nil {
		return err
	}

	ClassName := myos.Filename(destFile)

	// 生成Dart类定义
	classDef := fmt.Sprintf("class %s {\n", ClassName)

	// 生成字段定义
	fieldDefs := ""
	for _, field := range listField {
		fieldDefs += fmt.Sprintf("  %s %s;\n", field.Name, field.Type)
	}

	// 合并类定义和字段定义
	dartCode := classDef + fieldDefs + "}"

	return os.WriteFile(destFile, []byte(dartCode), 0777)
}
