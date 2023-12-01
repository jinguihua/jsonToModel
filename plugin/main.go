package main

import (
	"fmt"
	"github.com/any-call/gobase/util/myjson"
	"github.com/any-call/gobase/util/mystr"
	"github.com/jinguihua/jsonToModel/todart"
	"strings"
)

// 将json 转成 dart 模型
func JsonToDart(json string, className string) (ret string, err error) {
	if len(className) == 0 {
		return "", fmt.Errorf("入参类别不能为空")
	}

	list, err := myjson.ToOrderMap(json)
	if err != nil {
		return "", err
	}

	listClass := todart.FieldToClass(className, list)
	if listClass == nil || len(listClass) == 0 {
		return "", fmt.Errorf("covert err")
	}

	listDartInfo := make([]string, len(listClass)+1)
	listDartInfo[0] = "import 'dart:convert';"
	m := 1
	for i := (len(listClass) - 1); i >= 0; i-- {
		classInfo := listClass[i]
		dartModel, err := todart.NewTempModel(mystr.ToFirstUpper(classInfo.Name), classInfo.Fields)
		if err != nil {
			return "", err
		}

		dartInfo, err := dartModel.ToCode()
		if err != nil {
			return "", err
		}

		listDartInfo[m] = dartInfo
		m++
	}

	return strings.Join(listDartInfo, "\r\n\r\n"), nil
}
