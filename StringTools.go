package gaw

import (
	"html"
	"strings"
)

//String string
type String string

//FromString returns the String from a real string
func FromString(str string) String {
	return (String)(str)
}

//ToStringArray returns an array of real strings
func ToStringArray(arr []string) []String {
	sar := make([]String, len(arr))
	for i, v := range arr {
		sar[i] = (String)(v)
	}
	return sar
}

//ArrFromStringArray returns an array of real strings
func ArrFromStringArray(arr []String) []string {
	sar := make([]string, len(arr))
	for i, v := range arr {
		sar[i] = v.ToString()
	}
	return sar
}

//Trim trims a string
func (str *String) Trim() {
	*str = (String)(strings.Trim(str.ToString(), " "))
}

//TrimLeft trims a string on the left
func (str *String) TrimLeft() {
	*str = (String)(strings.TrimLeft(str.ToString(), " "))
}

//TrimRight trims a string on the left
func (str *String) TrimRight() {
	*str = (String)(strings.TrimRight(str.ToString(), " "))
}

//Relpace trims a string
func (str *String) Relpace(old, new String) {
	*str = (String)(strings.ReplaceAll(str.ToString(), old.ToString(), old.ToString()))
}

//SplitAfterN like strings.SplitAfterN
func (str *String) SplitAfterN(sep String, n int) []String {
	s := strings.SplitAfterN(str.ToString(), sep.ToString(), n)
	t := make([]String, len(s))
	for i, u := range s {
		t[i] = FromString(u)
	}
	return t
}

//SplitAfter like strings.SplitAfter
func (str *String) SplitAfter(sep String) []String {
	s := strings.SplitAfter(str.ToString(), sep.ToString())
	t := make([]String, len(s))
	for i, u := range s {
		t[i] = FromString(u)
	}
	return t
}

//Split like strings.Split
func (str *String) Split(sep String) []String {
	s := strings.Split(str.ToString(), sep.ToString())
	t := make([]String, len(s))
	for i, u := range s {
		t[i] = FromString(u)
	}
	return t
}

//ToUppercase like strings.ToUpper
func (str *String) ToUppercase() {
	*str = (String)(strings.ToUpper(str.ToString()))
}

//ToLowercase like strings.ToLower
func (str *String) ToLowercase() {
	*str = (String)(strings.ToLower(str.ToString()))
}

//ToString returns the real string
func (str *String) ToString() string {
	return string(*str)
}

//Length returns the length of the string
func (str *String) Length() int {
	return len(str.ToString())
}

//Count like strings.Count
func (str *String) Count(substr String) int {
	return strings.Count(str.ToString(), substr.ToString())
}

//BeginsWith like strings.HasPrefix
func (str *String) BeginsWith(prefix String) bool {
	return strings.HasPrefix(str.ToString(), prefix.ToString())
}

//EndsWith like strings.HasSuffix
func (str *String) EndsWith(prefix String) bool {
	return strings.HasSuffix(str.ToString(), prefix.ToString())
}

//EscapeSpecialChars avoid sqlInjection
func (str *String) EscapeSpecialChars() {
	*str = (String)(EscapeSpecialChars(str.ToString()))
}

//IsInArray returns true if the String array contains the given key
func (str *String) IsInArray(arr []String, args ...bool) bool {
	return IsInStringArray(str.ToString(), ArrFromStringArray(arr), args...)
}

//IsInStrArray returns true if the string array contains the given key
func (str *String) IsInStrArray(arr []string, args ...bool) bool {
	return IsInStringArray(str.ToString(), arr, args...)
}

//ToURL makes string to url
func (str *String) ToURL() (*URL, error) {
	return ParseURL(str.ToString())
}

//Contains like strings.Contains
func (str *String) Contains(strc String) bool {
	return strings.Contains(str.ToString(), strc.ToString())
}

//ContainsAny like strings.Contains
func (str *String) ContainsAny(chars String) bool {
	return strings.ContainsAny(str.ToString(), chars.ToString())
}

//IndexOf like strings.Index
func (str *String) IndexOf(substr String) int {
	return strings.Index(str.ToString(), substr.ToString())
}

//IndexOfAny like strings.IndexAny
func (str *String) IndexOfAny(substr String) int {
	return strings.IndexAny(str.ToString(), substr.ToString())
}

//LastIndexOf like strings.LastIndex
func (str *String) LastIndexOf(substr String) int {
	return strings.LastIndex(str.ToString(), substr.ToString())
}

//LastIndexOfAny like strings.IndexAny
func (str *String) LastIndexOfAny(substr String) int {
	return strings.LastIndexAny(str.ToString(), substr.ToString())
}

//Join like strings.Join
func (str *String) Join(str2 ...String) {
	*str = (String)(strings.Join(ArrFromStringArray(str2), str.ToString()))
}

//Fields like strings.Fields
func (str *String) Fields() []String {
	return ToStringArray(strings.Fields(str.ToString()))
}

//ToValidUTF8 like strings.ToValidUTF8
func (str *String) ToValidUTF8(relpacement String) {
	*str = FromString(strings.ToValidUTF8(str.ToString(), relpacement.ToString()))
}

//Map like strings.Map
func (str *String) Map(mapping func(rune) rune) String {
	return FromString(strings.Map(mapping, str.ToString()))
}

//Repeat like strings.Repeat
func (str *String) Repeat(count int) {
	*str = FromString(strings.Repeat(str.ToString(), count))
}

//Append appends text to the String str
func (str *String) Append(s String) {
	*str += s
}

//AppendIfNotEmpty appends s if str is not empty
func (str *String) AppendIfNotEmpty(s String) {
	if str.Length() > 0 {
		str.Append(s)
	}
}

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
