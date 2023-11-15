package todart

import (
	"fmt"
	"os"
	"testing"
	myTemp "text/template"
)

func Test_1(t *testing.T) {
	ts := &T{
		Add: func(i int) int {
			return i + 1
		},
	}

	tpl := `
		//只能使用 call 调用
		call field func Add :{{ call .ts.Add .y}}
		// 直接传入 .y 调用用
        call method func Sub: {{ .ts.Sub .y}}
		// and 测试 {{ and .m1 .m2 .m3}}
       // or 测试  {{ or .m1 .m2 .m3}}

	  //test 遍历
      {{ range $key,$value := .MapContent}} {{$key}} = {{$value}} {{end}}
     // if test
      {{ if .Name}} not empty {{else}}  empty {{end}}
`
	template1, err := myTemp.New("test").Parse(tpl)
	if err != nil {
		t.Error(err)
		return
	}

	if err = template1.Execute(os.Stdout, map[string]any{
		"y":          3,
		"ts":         ts,
		"m1":         1,
		"m2":         5,
		"m3":         3,
		"MapContent": map[int]string{1979: "jin", 1220: "gui", 532: "hua"},
		//"Name":       1,
	}); err != nil {
		t.Error(err)
		return
	}

	t.Logf("run ok")
}

type T struct {
	Add func(int) int
}

func (t *T) Sub(i int) int {
	fmt.Println("get argument i:", i)
	return i - 1
}

type Inventory struct {
	Material string
	Count    uint
}
