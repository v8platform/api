package v8tools

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/mash/go-tempfile-suffix"
	extraStrings "github.com/shomali11/util/strings"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"encoding/json"
	"fmt"
	"strings"
	"time"

	"../v8constants"
)

var (
	tempFiles []string
	tempDir   string = ИницализироватьВременныйКаталог()
)

func ВременныйКаталог() string {

	return ВременныйКаталогСПрефисом(v8constants.Prefix)
}

func ВременныйКаталогСПрефисом(pre string) string {

	t, err := ioutil.TempDir(tempDir, pre)
	if err != nil {
		panic(err)
	}
	tempFiles = append(tempFiles, t)
	return t
}

func ИницализироватьВременныйКаталог() string {

	userTmpDir := os.TempDir()

	tmpDir, err := ioutil.TempDir(userTmpDir, v8constants.Prefix)
	if err != nil {
		panic(err)
	}
	return tmpDir
}

func ЗначениеЗаполнено(Значение string) bool {
	return !extraStrings.IsEmpty(Значение)
}

func ПустаяСтрока(Значение string) bool {
	return extraStrings.IsEmpty(Значение)
}

func НовыйФайлИнформации() string {

	return НовыйВременныйФайл("", ".txt")
}

func НовыйВременныйФайл(p string, s string) string {

	f, err := tempfile.TempFileWithSuffix(tempDir, p, s)
	if err != nil {
		panic(err)
	}
	tempFiles = append(tempFiles, f.Name())

	return f.Name()
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
func IsNoExist(name string) (bool, error) {

	ok, err := Exists(name)
	return !ok, err
}

func ОчиститьВременныйКаталог() {

	for _, fileDir := range tempFiles {

		os.RemoveAll(fileDir)

	}

}

// Similar to ioutil.ReadFile() but decodes UTF-16.  Useful when
// reading data from MS-Windows systems that generate UTF-16BE files,
// but will do the right thing if other BOMs are found.
func ReadFileUTF16(filename string) ([]byte, error) {

	// Read the file into a []byte:
	raw, err := ioutil.ReadFile(filename)
	cs := DetectFileCharset(raw)

	if err != nil {
		return nil, err
	}

	var Endianness unicode.Endianness

	switch {
	case cs == Other:
		return raw, err
	case cs == Utf8withBOM:
		return raw[3:], err
	case cs == Utf16Be:
		Endianness = unicode.BigEndian
	case cs == Utf16Le:
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

type Charset byte

const (
	Utf8withBOM = Charset(iota)
	Utf16Be
	Utf16Le
	Other
)

func DetectFileCharset(data []byte) Charset {

	// Проверка на BOM
	if len(data) >= 3 {
		switch {
		case data[0] == 0xFF && data[1] == 0xFE:
			return Utf16Be
		case data[0] == 0xFE && data[1] == 0xFF:
			return Utf16Le
		case data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF:
			// wanna check special ascii codings here?
			return Utf8withBOM
		}
	}

	return Other
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewUID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// first create a type alias
type JsonBirthDate time.Time

// Add that to your struct
type Person struct {
	Name      string        `json:"name"`
	BirthDate JsonBirthDate `json:"birth_date"`
}

// imeplement Marshaler und Unmarshalere interface
func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	fmt.Print(t)
	//*j = JB(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Maybe a Format function for printing your date
func (j JsonBirthDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
