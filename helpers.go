package v8

import (
	"github.com/v8platform/errors"
	"io/ioutil"
	"strings"
)

////////////////////////////////////////////////////////
// Create InfoBases

func NewTempDir(dir, pattern string) string {

	t, _ := ioutil.TempDir(dir, pattern)

	return t

}

func NewTempFile(dir, pattern string) string {

	tempFile, _ := ioutil.TempFile(dir, pattern)

	defer tempFile.Close()

	return tempFile.Name()

}

var ErrorParseConnectingString = errors.BadConnectString.New("wrong connecting string")

func InfobaseFromConnectingStringOrNil(connectingString string) Infobase {
	ib, _ := InfobaseFromConnectingString(connectingString)
	return ib
}

func InfobaseFromConnectingString(connectingString string) (Infobase, error) {

	switch {
	case strings.Contains(connectingString, "Srvr="):
		return ServerInfobaseFromConnectingString(connectingString)
	case strings.Contains(connectingString, "File="):
		return FileInfobaseFromConnectingString(connectingString)
	default:
		return nil, ErrorParseConnectingString
	}

}

func FileInfobaseFromConnectingString(connectingString string) (FileInfoBase, error) {

	if strings.HasPrefix(connectingString, "/IBConnectionString ") {
		connectingString = strings.TrimPrefix(connectingString, "/IBConnectionString ")
	}

	var ib FileInfoBase
	values := strings.Split(connectingString, ";")

	for _, value := range values {

		if len(value) == 0 {
			continue
		}

		keyValue := strings.SplitN(value, "=", 2)

		if len(keyValue) != 2 {
			return FileInfoBase{}, errors.BadConnectString.New("wrong key/value count")
		}

		key := keyValue[0]
		val := keyValue[1]

		switch strings.ToLower(key) {

		case "file":
			ib.File = val
		case "locale":
			ib.Locale = val
		case "usr":
			ib.Usr = val
		case "pwd":
			ib.Pwd = val
		case "prmod":
			ib.Prmod = val == "1"
		case "licdstr":
			ib.LicDstr = val == "Y"
		case "zn":
			var err error
			ib.Zn, err = parseDatabaseSeparatorList(val)
			if err != nil {
				return FileInfoBase{}, err
			}
		}
	}

	if len(ib.File) == 0 {
		return FileInfoBase{}, errors.BadConnectString.New("wrong file connecting string")
	}

	return ib, nil

}

func ServerInfobaseFromConnectingString(connectingString string) (ServerInfoBase, error) {

	if strings.HasPrefix(connectingString, "/IBConnectionString ") {
		connectingString = strings.TrimPrefix(connectingString, "/IBConnectionString ")
	}

	var ib ServerInfoBase
	values := strings.Split(connectingString, ";")

	for _, value := range values {

		if len(value) == 0 {
			continue
		}

		keyValue := strings.SplitN(value, "=", 2)

		if len(keyValue) != 2 {
			return ServerInfoBase{}, errors.BadConnectString.New("wrong key/value count")
		}

		key := keyValue[0]
		val := keyValue[1]

		switch strings.ToLower(key) {

		case "srvr":
			ib.Srvr = val
		case "ref":
			ib.Ref = val
		case "usr":
			ib.Usr = val
		case "pwd":
			ib.Pwd = val
		case "licdstr":
			ib.LicDstr = val == "Y"
		case "prmod":
			ib.Prmod = val == "1"
		case "zn":
			var err error
			ib.Zn, err = parseDatabaseSeparatorList(val)
			if err != nil {
				return ServerInfoBase{}, err
			}
		}
	}

	if len(ib.Srvr) == 0 || len(ib.Ref) == 0 {
		return ServerInfoBase{}, errors.BadConnectString.New("wrong server connecting string")
	}

	return ib, nil
}

func parseDatabaseSeparatorList(stringValue string) (DatabaseSeparatorList, error) {
	// TODO Сделать парсер
	return DatabaseSeparatorList{}, nil
}
