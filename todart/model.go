package todart

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed temp.tpl
var tplFS string

type TempModel struct {
	*template.Template
	className string
	listField []Field
}

func NewTempModel(clsName string, fields []Field) (*TempModel, error) {
	tpl, err := template.New(clsName).Parse(tplFS)
	if err != nil {
		return nil, err
	}

	return &TempModel{
		Template:  tpl,
		className: clsName,
		listField: fields,
	}, nil
}

func (self *TempModel) ToCode() (string, error) {
	build := strings.Builder{}
	if err := self.Execute(&build, map[string]any{
		"ClassName": self.className,
		"genField": func() string {
			return self.genFieldDef()
		},
		"genConstruct": func() string {
			return self.genConstructFun()
		},
		"genCopyWithInputFun": func() string {
			return self.genCopyWithInputFun()
		},
		"genCopyWithRetFun": func() string {
			return self.genCopyWithRtnFun()
		},
		"genToStringFun": func() string {
			return self.genToStringFun()
		},
		"genOptEQFun": func() string {
			return self.genOptEquFun()
		},
		"genHashCodeFun": func() string {
			return self.genHashCodeFun()
		},
		"genFromMapFun": func() string {
			return self.genFromMapFun()
		},
		"genToMapFun": func() string {
			return self.genToMapFun()
		},
	}); err != nil {
		return "", err
	}

	return build.String(), nil
}

func (self *TempModel) genFieldDef() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	%s %s ;", field.Type, field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("%s %s ;\n", field.Type, field.Name)
		} else {
			ret += fmt.Sprintf("	%s %s ;\n", field.Type, field.Name)
		}
	}

	return ret
}

func (self *TempModel) genConstructFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	  this.%s ,", field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("this.%s ,\n", field.Name)
		} else {
			ret += fmt.Sprintf("	  this.%s ,\n", field.Name)
		}
	}

	return ret
}

func (self *TempModel) genCopyWithInputFun() string {
	var param string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			param += fmt.Sprintf("	  %s %s ,", field.Type, field.Name)
		} else if i == 0 {
			param += fmt.Sprintf("%s %s ,\n", field.Type, field.Name)
		} else {
			param += fmt.Sprintf("	  %s %s ,\n", field.Type, field.Name)
		}
	}

	return param
}

func (self *TempModel) genCopyWithRtnFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	    %s:%s ?? this.%s ,", field.Name, field.Name, field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("%s:%s ?? this.%s ,\n", field.Name, field.Name, field.Name)
		} else {
			ret += fmt.Sprintf("	    %s:%s ?? this.%s ,\n", field.Name, field.Name, field.Name)
		}
	}

	return ret
}

func (self *TempModel) genToStringFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) {
			ret += fmt.Sprintf("%s: $%s", field.Name, field.Name)
		} else {
			ret += fmt.Sprintf("%s: $%s,", field.Name, field.Name)
		}

	}

	return ret
}

func (self *TempModel) genOptEquFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	  	%s == other.%s", field.Name, field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("%s == other.%s &&\n", field.Name, field.Name)
		} else {
			ret += fmt.Sprintf("	  	%s == other.%s &&\n", field.Name, field.Name)
		}
	}

	return ret
}

func (self *TempModel) genHashCodeFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	    %s.hashCode", field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("%s.hashCode ^\n", field.Name)
		} else {
			ret += fmt.Sprintf("	 	  %s.hashCode ^\n", field.Name)
		}
	}

	return ret
}

func (self *TempModel) genFromMapFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	    %s: map['%s'] as  %s,", field.Name, field.Name, field.Type)
		} else if i == 0 {
			ret += fmt.Sprintf("%s: map['%s'] as  %s,\n", field.Name, field.Name, field.Type)
		} else {
			ret += fmt.Sprintf("	    %s: map['%s'] as  %s,\n", field.Name, field.Name, field.Type)
		}
	}

	return ret
}

func (self *TempModel) genToJsonFun() string {
	return fmt.Sprintf(`String toJson() => json.encode(toMap());`)
}

func (self *TempModel) genToMapFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("      '%s': this.%s,", field.Name, field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("'%s': this.%s,\n", field.Name, field.Name)
		} else {
			ret += fmt.Sprintf("      '%s': this.%s,\n", field.Name, field.Name)
		}

	}

	return ret
}
