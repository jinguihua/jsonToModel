package todart

import (
	"github.com/any-call/gobase/util/myjson"
	"github.com/any-call/gobase/util/myos"
	"github.com/any-call/gobase/util/mystr"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"
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
	listValue, err := myjson.ToOrderMap(jsonStr)
	if err != nil {
		return nil, err
	}

	ret = make([]Field, 0, len(listValue))
	for _, mapV := range listValue {
		mapV.Key = mystr.ToFirstLower(cases.Title(language.English, cases.NoLower).
			String(strings.Join(strings.Fields(mapV.Key), "")))

		field := Field{Name: mapV.Key, Value: mapV.Value}
		t := reflect.TypeOf(mapV.Value)
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.Type = "int?"
			break

		case reflect.Float32:
			{
				if mapV.Value == math.Trunc(mapV.Value.(float64)) {
					field.Type = "int?"
				} else {
					field.Type = "float?"
				}
			}
			break

		case reflect.Float64:
			{
				if mapV.Value == math.Trunc(mapV.Value.(float64)) {
					field.Type = "int?"
				} else {
					field.Type = "double?"
				}
			}
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
	dartModel, err := NewTempModel(mystr.ToTitle(ClassName), listField)
	if err != nil {
		return err
	}

	dartInfo, err := dartModel.ToCode()
	if err != nil {
		return err
	}

	return os.WriteFile(destFile, []byte(dartInfo), 0777)
}
