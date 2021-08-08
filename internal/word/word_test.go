package word

import "testing"

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	str:= "you_word"
 	t.Log(UnderscoreToUpperCamelCase(str))
}
