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

var ErrorParseConnectionString = errors.BadConnectString.New("wrong connection string")

func InfobaseFromConnectionStringOrNil(connectingString string) Infobase {
	ib, _ := InfobaseFromConnectionString(connectingString)
	return ib
}

func InfobaseFromConnectionString(connectingString string) (Infobase, error) {

	switch {
	case strings.Contains(connectingString, "Srvr="):
		return ServerInfobaseFromConnectionString(connectingString)
	case strings.Contains(connectingString, "File="):
		return FileInfobaseFromConnectionString(connectingString)
	default:
		return nil, ErrorParseConnectionString
	}

}

func FileInfobaseFromConnectionString(connectingString string) (FileInfoBase, error) {

	if strings.HasPrefix(connectingString, "/IBConnectionString ") {
		connectingString = strings.TrimPrefix(connectingString, "/IBConnectionString ")
	}

	var ib FileInfoBase
	values := strings.Split(connectingString, ";")

	for _, value := range values {

		if len(value) == 0 ||
			strings.HasPrefix(value, "/") ||
			strings.HasPrefix(value, "-") {
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
		return FileInfoBase{}, errors.BadConnectString.New("wrong file connection string")
	}

	return ib, nil

}

func ServerInfobaseFromConnectionString(connectingString string) (ServerInfoBase, error) {

	if strings.HasPrefix(connectingString, "/IBConnectionString ") {
		connectingString = strings.TrimPrefix(connectingString, "/IBConnectionString ")
	}

	var ib ServerInfoBase
	values := strings.Split(connectingString, ";")

	for _, value := range values {

		if len(value) == 0 ||
			strings.HasPrefix(value, "/") ||
			strings.HasPrefix(value, "-") {
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
		return ServerInfoBase{}, errors.BadConnectString.New("wrong server connection string")
	}

	return ib, nil
}

func parseDatabaseSeparatorList(stringValue string) (DatabaseSeparatorList, error) {
	// TODO Сделать парсер
	return DatabaseSeparatorList{}, nil
}
