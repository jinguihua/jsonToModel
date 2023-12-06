package main

import (
	"fmt"
	"github.com/any-call/gobase/util/myjson"
	"github.com/any-call/gobase/util/mystr"
	"github.com/jinguihua/jsonToModel/todart"
	"strings"
)

import "C"

var (
	JsonToDartResult int = 0
)

func init() {
	fmt.Println("enter init plugin")
}

func main() {
	//ret := JsonToDart(C.CString(jsonTemp), C.CString("Article"))
	//result := getLastResult()
	//fmt.Println("get run result :", result)
	//if result != 0 {
	//	fmt.Println("ret is :", C.GoString(ret))
	//}
}

//export JsonToDart
func JsonToDart(cjson *C.char, clsName *C.char) *C.char {
	json := C.GoString(cjson)
	className := C.GoString(clsName)
	if len(className) == 0 {
		JsonToDartResult = -1
		return C.CString("入参类别不能为空")
	}

	list, err := myjson.ToOrderMap(json)
	if err != nil {
		JsonToDartResult = -1
		return C.CString(err.Error())
	}

	listClass := todart.FieldToClass(className, list)
	if listClass == nil || len(listClass) == 0 {
		JsonToDartResult = -1
		return C.CString("convert err")
	}

	listDartInfo := make([]string, len(listClass)+1)
	listDartInfo[0] = "import 'dart:convert';"
	m := 1
	for i := (len(listClass) - 1); i >= 0; i-- {
		classInfo := listClass[i]
		dartModel, err := todart.NewTempModel(mystr.ToFirstUpper(classInfo.Name), classInfo.Fields)
		if err != nil {
			JsonToDartResult = -1
			return C.CString(err.Error())
		}

		dartInfo, err := dartModel.ToCode()
		if err != nil {
			JsonToDartResult = -1
			return C.CString(err.Error())
		}

		listDartInfo[m] = dartInfo
		m++
	}

	JsonToDartResult = 0
	return C.CString(strings.Join(listDartInfo, "\r\n\r\n"))
}

//export getLastResult
func getLastResult() int {
	return JsonToDartResult
}

const jsonTemp = `
{
    "ns_value": [
        "ns1.996dns.xyz",
        "ns2.996dns.xyz",
        "ns3.996dns.xyz",
        "ns4.996dns.xyz"
    ],
    "list_records:DomainRec": [
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 269,
            "host_record": "www",
            "record_type": "A",
            "record_value": "2.3.6.5",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 100000,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 270,
            "host_record": "www",
            "record_type": "A",
            "record_value": "6.5.2.3",
            "line_id": 9040,
            "line_area_map": {
                "3": {
                    "上海": true,
                    "北京": true,
                    "吉林": true
                },
                "4": {
                    "三亚": true,
                    "东城": true,
                    "南开": true,
                    "哈尔滨": true,
                    "嘉兴": true,
                    "大理": true,
                    "太原": true,
                    "安阳": true,
                    "徐州": true,
                    "徐汇": true,
                    "恩施": true,
                    "扬州": true,
                    "朝阳": true,
                    "枣庄": true,
                    "泰安": true,
                    "海口": true,
                    "湖州": true,
                    "益阳": true,
                    "船营": true,
                    "苏州": true,
                    "蚌埠": true,
                    "衢州": true,
                    "马鞍山": true,
                    "黄石": true
                },
                "5": {
                    "丰泽": true,
                    "信都": true,
                    "元宝": true,
                    "包河": true,
                    "双桥": true,
                    "城中": true,
                    "城关": true,
                    "天河": true,
                    "娄星": true,
                    "思明": true,
                    "昆山": true,
                    "桥西": true,
                    "武昌": true,
                    "武陵": true,
                    "江阴": true,
                    "沈河": true,
                    "浈江": true,
                    "海港": true,
                    "海陵": true,
                    "清城": true,
                    "玄武": true,
                    "盘龙": true,
                    "秀峰": true,
                    "章贡": true,
                    "芝罘": true,
                    "蓬江": true,
                    "西工": true,
                    "西市": true,
                    "西湖": true,
                    "西秀": true,
                    "西陵": true,
                    "越城": true,
                    "越秀": true,
                    "路北": true,
                    "鄂城": true,
                    "金平": true,
                    "铁东": true,
                    "锦江": true,
                    "雨花": true,
                    "鹤城": true,
                    "龙港": true
                }
            },
            "line_isp_map": {
                "美团": true
            },
            "is_define_line_id": false,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 2,
            "mx": 0,
            "ttl": 100000,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 323,
            "host_record": "ccc",
            "record_type": "A",
            "record_value": "2.3.2.5",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 20,
            "mx": 0,
            "ttl": 100000,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 324,
            "host_record": "ggg",
            "record_type": "A",
            "record_value": "43.227.112.40",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 100000,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 325,
            "host_record": "@",
            "record_type": "A",
            "record_value": "43.227.112.40",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 100000,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6543,
            "host_record": "www",
            "record_type": "A",
            "record_value": "96.2.3.2",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 100000,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6592,
            "host_record": "@",
            "record_type": "NS",
            "record_value": "ns1.996dns.xyz",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6593,
            "host_record": "@",
            "record_type": "NS",
            "record_value": "ns3.996dns.xyz",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6594,
            "host_record": "kjh",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6595,
            "host_record": "kpp",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6596,
            "host_record": "kpp2",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6597,
            "host_record": "dfg",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6598,
            "host_record": "qwe",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6599,
            "host_record": "qaz",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6600,
            "host_record": "mnb",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6601,
            "host_record": "mkl",
            "record_type": "A",
            "record_value": "103.82.140.147",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6604,
            "host_record": "eee",
            "record_type": "A",
            "record_value": "20.2.3.2",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6605,
            "host_record": "rrrr",
            "record_type": "A",
            "record_value": "52.3.2.1",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6606,
            "host_record": "dadada",
            "record_type": "A",
            "record_value": "23.2.2.2",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6607,
            "host_record": "rwtw",
            "record_type": "A",
            "record_value": "23.2.2.2",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6608,
            "host_record": "qwerg",
            "record_type": "A",
            "record_value": "45.2.3.2",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6609,
            "host_record": "tyu",
            "record_type": "A",
            "record_value": "78.23.3.2",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6610,
            "host_record": "wwrt",
            "record_type": "CNAME",
            "record_value": "wwe.cdc.com",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 180,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6611,
            "host_record": "wwwe",
            "record_type": "CNAME",
            "record_value": "uaydq.com",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 180,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 77,
            "domain_name": "v6dns.xyz",
            "record_id": 6612,
            "host_record": "qqqq",
            "record_type": "CNAME",
            "record_value": "adhuasia.com",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 180,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6485,
            "host_record": "a",
            "record_type": "A",
            "record_value": "2.3.6.5",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6490,
            "host_record": "php1",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6491,
            "host_record": "php2",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6492,
            "host_record": "node1",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6493,
            "host_record": "node2",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6494,
            "host_record": "go1",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6495,
            "host_record": "go2",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6496,
            "host_record": "other1",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6497,
            "host_record": "other2",
            "record_type": "A",
            "record_value": "43.227.112.157",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 599,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6588,
            "host_record": "@",
            "record_type": "NS",
            "record_value": "ns1.996dns.xyz",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 500,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6589,
            "host_record": "@",
            "record_type": "NS",
            "record_value": "ns2.996dns.xyz",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 500,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6602,
            "host_record": "test",
            "record_type": "A",
            "record_value": "1.1.1.1",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 60,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 98,
            "domain_name": "qqdns.xyz",
            "record_id": 6603,
            "host_record": "ttt",
            "record_type": "A",
            "record_value": "12.12.12.12",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 60,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 100,
            "domain_name": "v5dns.xyz",
            "record_id": 6590,
            "host_record": "@",
            "record_type": "NS",
            "record_value": "ns1.996dns.xyz",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 0,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 102,
            "domain_name": "v8dns.xyz",
            "record_id": 6584,
            "host_record": "@",
            "record_type": "NS",
            "record_value": "ns3.996dns.xyz",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 0,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 102,
            "domain_name": "v8dns.xyz",
            "record_id": 6585,
            "host_record": "@",
            "record_type": "NS",
            "record_value": "ns1.996dns.xyz",
            "line_id": 9065,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": true,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 1,
            "mx": 0,
            "ttl": 0,
            "status": 1,
            "cname": 0
        },
        {
            "domain_id": 3,
            "domain_name": "1.2.1",
            "record_id": 5,
            "host_record": "1.2.1.3",
            "record_type": "PTR",
            "record_value": "ceshi666.com",
            "line_id": 25,
            "line_area_map": {},
            "line_isp_map": {},
            "is_define_line_id": false,
            "is_custom_line": 0,
            "ip_sections": null,
            "weight": 0,
            "mx": 0,
            "ttl": 600,
            "status": 1,
            "cname": 0
        }
    ],
    "list_black_info": [
        {
            "domain": "oooommm.com",
            "type_name": "赌博",
            "style": 1,
            "cname": "xxcname.com"
        },
        {
            "domain": "domain5.com",
            "type_name": "诈骗",
            "style": 2,
            "cname": ""
        },
        {
            "domain": "ceshi.com",
            "type_name": "违法",
            "style": 2,
            "cname": ""
        },
        {
            "domain": "www.kkill.com",
            "type_name": "违法",
            "style": 1,
            "cname": "www.kkkkil.com"
        }
    ],
    "list_black_rmt_ip_info": [],
    "list_forward_domain": [
        {
            "domain": "ants22.com",
            "dns_servers": [
                "8.8.8.8"
            ]
        }
    ],
    "list_forward_section": [
        {
            "start_ip": "1.1.1.1",
            "end_ip": "255.255.255.255",
            "dns_servers": [
                "114.114.114.114",
                "8.8.8.8"
            ]
        }
    ],
    "list_qps_section": [
        {
            "domain_id": 102,
            "qps": 1000
        },
        {
            "domain_id": 101,
            "qps": 1000
        },
        {
            "domain_id": 98,
            "qps": 1000
        },
        {
            "domain_id": 100,
            "qps": 1000
        },
        {
            "domain_id": 103,
            "qps": 5500
        },
        {
            "domain_id": 77,
            "qps": 5500
        }
    ],
    "main_domain_ns_map": {
        "ddd.com": [
            "ns1.996dns.xyz",
            "ns2.996dns.xyz",
            "ns3.996dns.xyz"
        ],
        "qqdns.xyz": [
            "ns1.996dns.xyz",
            "ns2.996dns.xyz",
            "ns3.996dns.xyz"
        ],
        "v5dns.xyz": [
            "ns1.996dns.xyz",
            "ns2.996dns.xyz"
        ],
        "v6dns.xyz": [
            "ns1.996dns.xyz",
            "ns2.996dns.xyz",
            "ns3.996dns.xyz"
        ],
        "v8dns.xyz": [
            "ns1.996dns.xyz",
            "ns2.996dns.xyz",
            "ns3.996dns.xyz"
        ],
        "vv5dns.xyz": [
            "ns1.996dns.xyz",
            "ns2.996dns.xyz",
            "ns3.996dns.xyz"
        ]
    },
    "site_ip_domain_map": {
        "43.227.112.128": [
            "v6dns.xyz",
            "qqdns.xyz",
            "v5dns.xyz",
            "vv5dns.xyz",
            "v8dns.xyz",
            "ddd.com"
        ],
        "43.227.112.149": [
            "v6dns.xyz",
            "qqdns.xyz",
            "v5dns.xyz",
            "vv5dns.xyz",
            "v8dns.xyz",
            "ddd.com"
        ],
        "43.227.112.170": [
            "v6dns.xyz",
            "qqdns.xyz",
            "vv5dns.xyz",
            "v8dns.xyz",
            "ddd.com"
        ]
    },
    "black_setting": {
        "site_qps": 200000,
        "domain_rec_qps": 200000
    },
    "site_warn": {
        "129.3.2.3": 200000,
        "43.227.112.128": 70000,
        "43.227.112.149": 7000000,
        "43.227.112.170": 500000
    }
}
`
