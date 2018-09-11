package chunk

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func decodeUTF16(data []byte) string {
	buf := bytes.NewBuffer(data)
	transformer := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	decoder := transformer.NewDecoder()
	r := transform.NewReader(buf, unicode.BOMOverride(decoder))
	s, _ := ioutil.ReadAll(r)
	return string(s)
}
