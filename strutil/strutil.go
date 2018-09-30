package strutil

import (
	"bytes"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// DecodeUTF16 decodes UTF-16 byte slice to a UTF-8 string.
func DecodeUTF16(data []byte) string {
	buf := bytes.NewBuffer(data)
	transformer := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	decoder := transformer.NewDecoder()
	r := transform.NewReader(buf, unicode.BOMOverride(decoder))
	s, _ := ioutil.ReadAll(r)
	return string(s)
}

// TrimNull trims any null bytes from the start and end of the given string.
func TrimNull(str string) string {
	return strings.Trim(str, string(0))
}
