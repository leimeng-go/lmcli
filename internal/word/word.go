package word

import (
	"strings"
	"unicode"
)

func ToLower(s string) string {
	return strings.ToLower(s)
}
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//UnderscoreToLowerCamelCase 小驼峰
func UnderscoreToLowerCamelCase(s string) string {
	result := ""
	array := strings.Split(s, "_")
	for i, v := range array {
		if i > 0 {
			v = strings.Title(v)
		}
		result += v
	}
	return result
}

//UnderscoreToLowerCamelCaseV1 小驼峰V1
func UnderscoreToLowerCamelCaseV1(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//UnderscoreToUpperCamelCase 大驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
