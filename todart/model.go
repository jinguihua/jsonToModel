package todart

import (
	"fmt"
)

const (
	keyClassName      = "${ClassName}"
	keyFieldsDef      = "${FieldDefine}"
	keyConstructParam = "$Construct param"
	KeyCopyWithParam  = "$copyWith param"
	KeyCopyWithReturn = "$copyWith return"
)

type TempModel struct {
	className string
	listField []Field
}

func NewTempModel(clsName string, fields []Field) *TempModel {
	return &TempModel{
		className: clsName,
		listField: fields,
	}
}

func (self *TempModel) ToCode() string {
	return fmt.Sprintf(template,
		self.className,
		self.genFieldDef(),
		self.genConstructFun(),
		self.genCopyWithFun(),
		self.genToStringFun(),
		self.genOptEquFun(),
		self.genHashCodeFun(),
		self.genFromMapFun(),
		self.genToMapFun(),
	)
}

func (self *TempModel) genFieldDef() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	%s %s ;\n", field.Type, field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("%s %s ;\n", field.Type, field.Name)
		} else {
			ret += fmt.Sprintf("	%s %s ;\n", field.Type, field.Name)
		}
	}

	return fmt.Sprintf(`%s
`, ret)
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

	return fmt.Sprintf(`%s({
	  %s	
   });
`, self.className, ret)
}

func (self *TempModel) genCopyWithFun() string {
	var param, ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			param += fmt.Sprintf("	  %s %s ,", field.Type, field.Name)
			ret += fmt.Sprintf("	    %s:%s ?? this.%s ,", field.Name, field.Name, field.Name)
		} else if i == 0 {
			param += fmt.Sprintf("	%s %s ,\n", field.Type, field.Name)
			ret += fmt.Sprintf("	  %s:%s ?? this.%s ,\n", field.Name, field.Name, field.Name)
		} else {
			param += fmt.Sprintf("	  %s %s ,\n", field.Type, field.Name)
			ret += fmt.Sprintf("	    %s:%s ?? this.%s ,\n", field.Name, field.Name, field.Name)
		}
	}

	return fmt.Sprintf(` %s copyWith({
	%s	
  }) {
    return new %s(
	%s
    );
  }`, self.className, param, self.className, ret)
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

	return fmt.Sprintf(`@override
  String toString() {
	  return '%s{%s}';
  }
`, self.className, ret)
}

func (self *TempModel) genOptEquFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	  	%s == other.%s", field.Name, field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("	%s == other.%s && \n", field.Name, field.Name)
		} else {
			ret += fmt.Sprintf("	  	%s == other.%s && \n", field.Name, field.Name)
		}
	}

	return fmt.Sprintf(`@override
  bool operator ==(Object other) =>
      identical(this, other) ||
      (other is %s && 
      %s) ;
`, self.className, ret)
}

func (self *TempModel) genHashCodeFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	  %s.hashCode ^", field.Name)
		} else if i == 0 {
			ret += fmt.Sprintf("%s.hashCode ^\n", field.Name)
		} else {
			ret += fmt.Sprintf("	  %s.hashCode ^\n", field.Name)
		}
	}

	return fmt.Sprintf(`@override
int get hashCode =>
    %s ;
`, ret)
}

func (self *TempModel) genFromMapFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	%s: map['%s'] as  %s,", field.Name, field.Name, field.Type)
		} else {
			ret += fmt.Sprintf("	%s: map['%s'] as  %s, \n", field.Name, field.Name, field.Type)
		}

	}

	return fmt.Sprintf(`
  factory %s.fromMap(Map<String?, dynamic> map) {
    return new %s(
       %s 
	);
  }
`, self.className, self.className, ret)
}

func (self *TempModel) genToMapFun() string {
	var ret string
	for i, field := range self.listField {
		if (i + 1) == len(self.listField) { //最后一个
			ret += fmt.Sprintf("	'%s': this.%s,", field.Name, field.Name)
		} else {
			ret += fmt.Sprintf("	'%s': this.%s, \n", field.Name, field.Name)
		}

	}

	return fmt.Sprintf(`
  Map<String, dynamic> toMap() {
    // ignore: unnecessary_cast
    return {
       %s
     } as Map<String, dynamic>;
  }
`, ret)
}

const template = `
class %s {
  //field define 
  %s

 //construct function 
  %s

 //copyWith function
 %s

//toString function
 %s

//operation == function
 %s
 
//get hasCode function
%s

//fromMap function
%s

//Map function
%s

}
`
