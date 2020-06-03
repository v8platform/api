package cmd

import (
	"bytes"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type charset byte

const (
	utf8withBOM = charset(iota)
	utf16Be
	utf16Le
	other
)

func detectFileCharset(data []byte) charset {

	// Проверка на BOM
	if len(data) >= 3 {
		switch {
		case data[0] == 0xFF && data[1] == 0xFE:
			return utf16Be
		case data[0] == 0xFE && data[1] == 0xFF:
			return utf16Le
		case data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF:
			// wanna Check special ascii codings here?
			return utf8withBOM
		}
	}

	return other
}

// Similar to ioutil.ReadFile() but decodes UTF-16.  Useful when
func readV8file(filename string) ([]byte, error) {

	// Reader the file into a []byte:
	raw, err := ioutil.ReadFile(filename)
	cs := detectFileCharset(raw)

	if err != nil {
		return nil, err
	}

	var Endianness unicode.Endianness

	switch {
	case cs == other:

		// Make a Reader that uses utf16bom:
		unicodeReader := transform.NewReader(bytes.NewReader(raw), charmap.Windows1251.NewDecoder())

		// decode and print:
		decoded, err := ioutil.ReadAll(unicodeReader)
		return decoded, err

	case cs == utf8withBOM:
		return raw[3:], err
	case cs == utf16Be:
		Endianness = unicode.BigEndian
	case cs == utf16Le:
		Endianness = unicode.LittleEndian
	}

	// Make an tranformer that converts MS-Win default to UTF8:
	win16be := unicode.UTF16(Endianness, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)

	// decode and print:
	decoded, err := ioutil.ReadAll(unicodeReader)
	return decoded, err

}
