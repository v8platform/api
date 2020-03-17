package sshclient

type Agent interface {
	Exec(Cmd AgentCommand, opts ...execOption) (res []Respond, err error)

	CopyFileTo(src, dest string) error
	CopyFileFrom(src, dest string) error

	//Команды группы common отвечают за общие операции. В состав группы входят следующие команды:
	//connect-ib ‑ выполнить подключение к информационной базе, параметры которой указаны при старте режима агента.
	Connect(opts ...execOption) (err error)

	//disconnect-ib ‑ выполнить отключение от информационной базы, подключение к которой ранее выполнялось с помощью команды connect-ib.
	Disconnect(opts ...execOption) (err error)

	//shutdown ‑ завершить работу конфигуратора в режиме агента.
	Shutdown(opts ...execOption) (err error)

	// options
	Options(opts ...execOption) (confOpts ConfigurationOptions, err error)
	SetOptions(confOpts ConfigurationOptions, opts ...execOption) error

	// Configuration support
	DisableCfgSupport(opts ...execOption) error

	// Configuration
	DumpCfgToFiles(dir string, force bool, opts ...execOption) error
	LoadCfgFromFiles(dir string, updateConfigDumpInfo bool, opts ...execOption) error

	DumpCfg(file string, opts ...execOption) error
	LoadCfg(file string, opts ...execOption) error

	DumpExtensionCfg(ext string, file string, opts ...execOption) error
	LoadExtensionCfg(ext string, file string, opts ...execOption) error

	DumpExtensionToFiles(ext string, dir string, force bool, opts ...execOption) error
	LoadExtensionFromFiles(ext string, dir string, updateConfigDumpInfo bool, opts ...execOption) error
	DumpAllExtensionsToFiles(dir string, force bool, opts ...execOption) error
	LoadAllExtensionsFromFiles(dir string, updateConfigDumpInfo bool, opts ...execOption) error

	// update
	UpdateDbCfg(server bool, opts ...execOption) error
	UpdateDbExtension(extension string, server bool, opts ...execOption) error
	StartBackgroundUpdateDBCfg(opts ...execOption) error
	StopBackgroundUpdateDBCfg(opts ...execOption) error
	FinishBackgroundUpdateDBCfg(opts ...execOption) error
	ResumeBackgroundUpdateDBCfg(opts ...execOption) error

	// Infobase
	IBDataSeparationList(opts ...execOption) (DataSeparationList, error)
	DebugInfo(opts ...execOption) (info DebugInfo, err error)
	DumpIB(file string, opts ...execOption) (err error)
	RestoreIB(file string, opts ...execOption) (err error)
	EraseData(opts ...execOption) (err error)

	//Extensions
	CreateExtension(name, prefix string, synonym string, purpose ExtensionPurposeType, opts ...execOption) error
	DeleteExtension(name string, opts ...execOption) error
	DeleteAllExtensions(opts ...execOption) error
	GetExtensionProperties(name string, opts ...execOption) (ExtensionProperties, error)
	GetAllExtensionsProperties(opts ...execOption) ([]ExtensionProperties, error)
	SetExtensionProperties(props ExtensionProperties, opts ...execOption) error
}

type DebugInfo struct {

	//  enabled ‑ признак включения отладки.
	Enable bool `json:"enable"`

	//  protocol ‑ протокол отладки: tcp или http.
	Protocol string `json:"protocol"`

	//  server-address ‑ адрес сервера отладки для данной информационной базы.
	ServerAddress string `json:"server-address"`
}

type ExtensionProperties struct {
	Extension                 string
	Active                    ExtensionPropertiesBoolType
	SafeMode                  ExtensionPropertiesBoolType
	SecurityProfileName       string
	UnsafeActionProtection    ExtensionPropertiesBoolType
	UsedInDistributedInfobase ExtensionPropertiesBoolType
	Scope                     ExtensionPropertiesScopeType
}

//Данная команда позволяет получить значения параметров. Для команды доступны следующие параметры:
type ConfigurationOptions struct {

	//  --output-format ‑ позволяет указать формат вывода результата работы команд:
	//  text ‑ команды возвращают результат в текстовом формате.
	//  json ‑ команды возвращают результат в формате JSON-сообщений.
	OutputFormat OptionsOutputFormatType `json:"output-format"`

	//  --show-prompt ‑ позволяет управлять наличием приглашения командной строки designer>:
	//  yes ‑ в командной строке есть приглашение;
	//  no ‑ в командной строке нет приглашения.
	ShowPrompt bool `json:"show-prompt"`

	//  --notify-progress ‑ позволяет получить информацию об отображении прогресса выполнения команды.
	NotifyProgress bool `json:"notify-progress"`

	//  --notify-progress-interval ‑ позволяет получить интервал времени, через который обновляется информация о прогрессе.
	NotifyProgressInterval int `json:"notify-progress-interval"`
}
