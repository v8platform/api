package infobase

import (
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

func (ib *Server) Path() string {
	return ib.Srvr + "\\" + ib.Ref
}

func (ib *Server) Auth(user, password string) {
	if len(user) == 0 {
		return
	}
	ib.Usr = user
	ib.Pwd = password
}

func (ib *Server) DatabaseSeparator(list DatabaseSeparatorList) {
	ib.Zn = list
}

func (ib *Server) ConnectionString() string {

	v, _ := marshaler.Marshal(ib)
	connString := strings.Join(v, ";")
	return "/IBConnectionString " + connString
}
