package todart

import (
	"fmt"
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

type (
	Field struct {
		Key      string
		Name     string
		Type     string
		ItemName string //基于约定:itemName
		Value    any
	}

	Class struct {
		Name   string
		Fields []Field
	}
)

// 针对map ,list 获取children 的类型
func (self *Field) GetItemValueType() string {
	if list, ok := self.Value.(myjson.OrderArray); ok {
		for _, v := range list {
			t := reflect.TypeOf(v)
			switch t.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				return "int"
			case reflect.Float32:
				{
					if v == math.Trunc(v.(float64)) {
						return "int"
					} else {
						return "float"
					}
				}
			case reflect.String:
				return "String"

			case reflect.Bool:
				return "bool"

			case reflect.Map, reflect.Struct:
				if self.ItemName != "" {
					return self.ItemName
				}
				return "Map"

			case reflect.Slice, reflect.Array:
				return "List"
			}
		}
	}

	return ""
}

func fieldToClass(clsName string, listField []myjson.FieldValue) []Class {
	var listClass = make([]Class, 0)

	if listField == nil || len(listField) == 0 {
		return listClass
	}

	ret := Class{
		Name:   clsName,
		Fields: make([]Field, 0, len(listField)),
	}
	for _, mapV := range listField {
		tmpName := mapV.Key
		listName := strings.Split(tmpName, ":")
		var parName, itemName string
		if len(listName) == 2 {
			parName = listName[0]
			itemName = listName[1]
		} else {
			parName = tmpName
		}

		parName = strings.Join(strings.Fields(parName), "")
		itemName = strings.Join(strings.Fields(itemName), "")

		parName = mystr.ToCamel(parName)
		parName = mystr.ToFirstLower(cases.Title(language.English, cases.NoLower).String(parName))

		field := Field{Key: mapV.Key, Name: parName, Value: mapV.Value, ItemName: itemName}
		if len(listName) == 2 {
			field.Key = listName[0]
		}

		t := reflect.TypeOf(mapV.Value)
		if t == nil {
			continue
		}
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

		case reflect.Map, reflect.Struct:
			field.Type = "Map?"
			break

		case reflect.Slice, reflect.Array:
			{
				itemType := field.GetItemValueType()
				if itemType == "" { //说明没有下一级
					field.Type = "List?"
				} else {
					field.Type = fmt.Sprintf("List<%s>?", itemType)
					if field.ItemName != "" { //说明标识有下一级的类名
						if list, ok := mapV.Value.(myjson.OrderArray); ok && len(list) > 0 {
							if orderMap, ok := list[0].(myjson.OrderMap); ok {
								tempClass := fieldToClass(field.ItemName, orderMap.Fields())
								if tempClass != nil && len(tempClass) > 0 {
									field.Type = fmt.Sprintf("List<%s>?", field.ItemName)
									listClass = append(listClass, tempClass...)
								}
							}
						}
					}
				}
			}

			break

		default:
			field.Type = "Object?"
			break
		}

		ret.Fields = append(ret.Fields, field)
	}

	listClass = append(listClass, ret)
	return listClass
}

func GenModel(destFile string, jsonStr string) error {
	list, err := myjson.ToOrderMap(jsonStr)
	if err != nil {
		return err
	}

	fileNameSlice := strings.Split(myos.Filename(destFile), ".")
	ClassName := fileNameSlice[0]

	listClass := fieldToClass(ClassName, list)
	if listClass == nil || len(listClass) == 0 {
		return fmt.Errorf("covert err")
	}

	listDartInfo := make([]string, len(listClass)+1)
	listDartInfo[0] = "import 'dart:convert';"
	m := 1
	for i := (len(listClass) - 1); i >= 0; i-- {
		classInfo := listClass[i]
		dartModel, err := NewTempModel(mystr.ToFirstUpper(classInfo.Name), classInfo.Fields)
		if err != nil {
			return err
		}

		dartInfo, err := dartModel.ToCode()
		if err != nil {
			return err
		}

		listDartInfo[m] = dartInfo
		m++
	}

	return os.WriteFile(destFile, []byte(strings.Join(listDartInfo, "\r\n\r\n")), 0777)
}
