package goawesomehelper

import (
	"html"
	"strings"
)

//EscapeSpecialChars avoid sqlInjection
func EscapeSpecialChars(inp string) string {
	if len(inp) == 0 {
		return ""
	}
	toReplace := []string{"'", "`", "\""}
	for _, i := range toReplace {
		inp = strings.ReplaceAll(inp, i, "")
	}
	return html.EscapeString(inp)
}

//IsInStringArray returns true if the array contains the given key
func IsInStringArray(str string, arr []string, args ...bool) bool {
	var trim bool
	if len(args) > 0 {
		trim = args[0]
	}
	if trim {
		str = strings.Trim(str, " ")
	}
	for _, s := range arr {
		z := s
		if trim {
			z = strings.Trim(z, " ")
		}
		if z == str {
			return true
		}
	}
	return false
}
