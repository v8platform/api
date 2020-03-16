package sshclient

type OutputFormatType string

const (
	OutputFormatText OutputFormatType = "text"
	OutputFormatJson                  = "json"
)

type DebugInfo struct {

	//  enabled ‑ признак включения отладки.
	Enable bool `json:"enable"`

	//  protocol ‑ протокол отладки: tcp или http.
	Protocol string `json:"protocol"`

	//  server-address ‑ адрес сервера отладки для данной информационной базы.
	ServerAddress string `json:"server-address"`
}

//Данная команда позволяет получить значения параметров. Для команды доступны следующие параметры:
type ConfigurationOptions struct {

	//  --output-format ‑ позволяет указать формат вывода результата работы команд:
	//  text ‑ команды возвращают результат в текстовом формате.
	//  json ‑ команды возвращают результат в формате JSON-сообщений.
	OutputFormat OutputFormatType `json:"output-format"`

	//  --show-prompt ‑ позволяет управлять наличием приглашения командной строки designer>:
	//  yes ‑ в командной строке есть приглашение;
	//  no ‑ в командной строке нет приглашения.
	ShowPrompt bool `json:"show-prompt"`

	//  --notify-progress ‑ позволяет получить информацию об отображении прогресса выполнения команды.
	NotifyProgress bool `json:"notify-progress"`

	//  --notify-progress-interval ‑ позволяет получить интервал времени, через который обновляется информация о прогрессе.
	NotifyProgressInterval int `json:"notify-progress-interval"`
}
