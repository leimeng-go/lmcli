package word

import (
	"strings"
)

func ToLower(s string)string{
	return strings.ToLower(s)
}
func ToUpper(s string)string{
	return strings.ToUpper(s)
}
func UnderscoreToUpperCamelCase(s string)string{
	result:=""
	array:=strings.Split(s, "_")
	for i,v:=range array{
		if i>0{
          v=strings.Title(v)
		}
		result+=v
	}
	return result
}