package word

import "testing"

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	str := "you_word"
	t.Logf("小驼峰: %s", UnderscoreToLowerCamelCase(str))
	t.Logf("小驼峰V1: %s", UnderscoreToLowerCamelCaseV1(str))
	t.Logf("大驼峰: %s", UnderscoreToUpperCamelCase(str))
	t.Logf("小驼峰转下滑: %s", CamelCaseToUnderscore(UnderscoreToLowerCamelCaseV1(str)))
	t.Logf("大驼峰转下滑: %s", CamelCaseToUnderscore(UnderscoreToUpperCamelCase(str)))
}
