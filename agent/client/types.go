package sshclient

type Agent interface {
	Exec(Cmd AgentCommand, opts ...execOption) (res []Respond, err error)

	CopyFileTo(src, dest string) error
	CopyFileFrom(src, dest string) error

	//Команды группы common отвечают за общие операции. В состав группы входят следующие команды:
	//connect-ib ‑ выполнить подключение к информационной базе, параметры которой указаны при старте режима агента.
	Connect() (err error)

	//disconnect-ib ‑ выполнить отключение от информационной базы, подключение к которой ранее выполнялось с помощью команды connect-ib.
	Disconnect() (err error)

	//shutdown ‑ завершить работу конфигуратора в режиме агента.
	Shutdown() (err error)

	// options
	Options() (opts ConfigurationOptions, err error)
	SetOptions(opts ConfigurationOptions) error

	// Configuration support
	DisableCfgSupport() error

	// Configuration
	DumpCfgToFiles(dir string, force bool) error
	LoadCfgFromFiles(dir string, updateConfigDumpInfo bool) error

	DumpCfg(file string) error
	LoadCfg(file string) error

	DumpExtensionCfg(ext string, file string) error
	LoadExtensionCfg(ext string, file string) error

	DumpExtensionToFiles(ext string, dir string, force bool) error
	LoadExtensionFromFiles(ext string, dir string, updateConfigDumpInfo bool) error
	DumpAllExtensionsToFiles(dir string, force bool) error
	LoadAllExtensionsFromFiles(dir string, updateConfigDumpInfo bool) error

	// update
	UpdateDbCfg(server bool) error
	UpdateDbExtension(extension string, server bool) error
	StartBackgroundUpdateDBCfg() error
	StopBackgroundUpdateDBCfg() error
	FinishBackgroundUpdateDBCfg() error
	ResumeBackgroundUpdateDBCfg() error

	// Infobase
	IBDataSeparationList() (DataSeparationList, error)
	DebugInfo() (info DebugInfo, err error)
	DumpIB(file string) (err error)
	RestoreIB(file string) (err error)
	EraseData() (err error)

	//Extensions
	CreateExtension(name, prefix string, synonym string, purpose ExtensionPurposeType) error
	DeleteExtension(name string) error
	DeleteAllExtensions() error
	GetExtensionProperties(name string) (ExtensionProperties, error)
	GetAllExtensionsProperties() ([]ExtensionProperties, error)
	SetExtensionProperties(props ExtensionProperties) error
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
