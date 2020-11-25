package infobase

import (
	"github.com/v8platform/errors"
	"github.com/v8platform/marshaler"
	"strings"
)

var _ Infobase = (*Server)(nil)

type Server struct {
	Common `v8:",inherit" json:"infobase"`

	//имя сервера «1С:Предприятия» в формате: [<протокол>://]<адрес>[:<порт>], где:
	//<протокол> – не обязателен, поддерживается только протокол TCP,
	//<адрес> – имя сервера или IP-адрес сервера в форматах IPv4 или IPv6,
	//<порт> – не обязателен, порт главного менеджера кластера, по умолчанию равен 1541.
	Srvr string `v8:"Srvr, equal_sep" json:"srvr"`

	//имя информационной базы на сервере "1С:Предприятия";
	Ref string `v8:"Ref, equal_sep, quotes" json:"ref"`
}

func (ib Server) Path() string {
	return ib.Srvr + "\\" + ib.Ref
}

func (ib Server) CommonValues() Common {

	return ib.Common

}

func (ib Server) withCommonValues(values Common) Infobase {

	newIb := ib
	newIb.Common = values
	return newIb
}

func (ib Server) ConnectionString() string {

	v, _ := marshaler.Marshal(ib)
	connString := strings.Join(v, ";")
	return "/IBConnectionString " + connString
}

func (ib Server) Parse(connectingString string) error {

	if strings.HasPrefix(connectingString, "/IBConnectionString ") {
		connectingString = strings.TrimPrefix(connectingString, "/IBConnectionString ")
	}

	ibPtr := &ib

	values := strings.Split(connectingString, ";")

	for _, value := range values {

		if len(value) == 0 ||
			strings.HasPrefix(value, "/") ||
			strings.HasPrefix(value, "-") {
			continue
		}

		keyValue := strings.SplitN(value, "=", 2)

		if len(keyValue) != 2 {
			return errors.BadConnectString.New("wrong key/value count")
		}

		key := keyValue[0]
		val := keyValue[1]

		switch strings.ToLower(key) {

		case "srvr":
			ibPtr.Srvr = val
		case "ref":
			ibPtr.Ref = val
		case "usr":
			ibPtr.Usr = val
		case "pwd":
			ibPtr.Pwd = val
		case "licdstr":
			ibPtr.LicDstr = val == "Y"
		case "prmod":
			ibPtr.Prmod = val == "1"
		case "zn":
			var err error
			ibPtr.Zn, err = ParseDatabaseSeparatorList(val)
			if err != nil {
				return err
			}
		}
	}

	if len(ibPtr.Srvr) == 0 || len(ibPtr.Ref) == 0 {
		return errors.BadConnectString.New("wrong server connection string")
	}

	return nil

}
