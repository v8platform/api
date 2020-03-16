package sshclient

import (
	"fmt"
	"strconv"
	"strings"
)

type AgentCommand interface {
	Command() string
	Args() []string
}

type cmdCommon struct{}

func (c cmdCommon) Command() string {
	return "common"
}

func (cmdCommon) Args() (v []string) {
	return
}

type CommonConnectInfobase struct {
	cmdCommon
}

func (c CommonConnectInfobase) Command() string {
	return fmt.Sprintf("%s %s", c.cmdCommon.Command(), "connect-ib")
}

type CommonDisconnectInfobase struct {
	cmdCommon
}

func (c CommonDisconnectInfobase) Command() string {
	return fmt.Sprintf("%s %s", c.cmdCommon.Command(), "disconnect-ib")
}

type CommonShutdown struct {
	cmdCommon
}

func (c CommonShutdown) Command() string {
	return fmt.Sprintf("%s %s", c.cmdCommon.Command(), "shutdown")
}

type cmdOptions struct{}

func (c cmdOptions) Command() string {
	return "options"
}

func (cmdOptions) Args() (v []string) {
	return
}

type OptionsBoolType string

func (t OptionsBoolType) String() string {
	return string(t)
}

const (
	OptionsBoolYes    OptionsBoolType = "yes"
	OptionsBoolNo     OptionsBoolType = "no"
	OptionsBoolNotSet                 = ""
)

type OptionsOutputFormatType string

const (
	OptionsOutputFormatText OptionsOutputFormatType = "text"
	OptionsOutputFormatJson                         = "json"
)

type OptionsList struct {
	cmdOptions
}

func (c OptionsList) Command() string {
	return fmt.Sprintf("%s %s", c.cmdOptions.Command(), "list")
}

func (t OptionsOutputFormatType) String() string {
	return string(t)
}

type SetOptions struct {
	cmdOptions
	//  --output-format ‑ позволяет указать формат вывода результата работы команд:
	//  text ‑ команды возвращают результат в текстовом формате.
	//  json ‑ команды возвращают результат в формате JSON-сообщений.
	OutputFormat OptionsOutputFormatType `json:"output-format"`
	//  --show-prompt ‑ позволяет управлять наличием приглашения командной строки designer>:
	//  yes ‑ в командной строке есть приглашение;
	//  no ‑ в командной строке нет приглашения.
	ShowPrompt OptionsBoolType `json:"show-prompt"`
	//  --notify-progress ‑ позволяет получить информацию об отображении прогресса выполнения команды.
	NotifyProgress bool `json:"notify-progress"`
	//  --notify-progress-interval ‑ позволяет получить интервал времени, через который обновляется информация о прогрессе.
	NotifyProgressInterval int `json:"notify-progress-interval"`
}

func (c SetOptions) Command() string {
	return fmt.Sprintf("%s %s", c.cmdOptions.Command(), "set")
}

func (c SetOptions) Args() []string {

	var v []string

	if len(c.OutputFormat) > 0 {
		v = append(v, "--output-format", c.OutputFormat.String())
	}

	if len(c.ShowPrompt) > 0 {
		v = append(v, "--show-prompt", c.ShowPrompt.String())
	}

	if c.NotifyProgress {
		v = append(v, "--notify-progress")
	}

	if c.NotifyProgressInterval > 0 {
		v = append(v, "--notify-progress-interval", strconv.FormatInt(int64(c.NotifyProgressInterval), 10))
	}

	return v
}

//Команды группы config отвечают за команды редактирования конфигурации.
type cmdConfig struct{}

func (c cmdConfig) Command() string {
	return "config"
}

//dump-config-to-files
//Команда позволяет выполнить выгрузку конфигурации в xml-файлы.
type DumpCfgToFiles struct {
	cmdConfig
	//--dir <путь> ‑ содержит путь к каталогу, в который будет выгружена конфигурация. Параметр является обязательным.
	Dir string `json:"dir"`
	//--extension <имя расширения> ‑ содержит имя расширения, которое будет выгружено в файл.
	Extension string `json:"extension"`
	//--all-extensions ‑ если параметр указан, то в файлы будут выгружены все расширения конфигурации.
	AllExtension bool `json:"all-extension"`
	//--format [hierarchical|plain] ‑ определяет формат выгрузки. По умолчанию используется иерархический формат выгрузки (hierarchical).
	Format string
	//--update ‑ обновить существующую выгрузку. В этом случае будут выгружены только те объекты, версии которых отличаются от версий, указанных в файле ConfigDumpInfo.xml.
	Update bool
	//--force ‑ выполнить полную выгрузку, если при попытке обновления выгрузки (параметр update) выяснилось, что текущая версия формата выгрузки не совпадает с версией формата выгрузки, которая указана в файле ConfigDumpInfo.xml.
	Force bool
	//--get-сhanges <путь> ‑ сформировать файл, который содержит изменения между текущей и указанной выгрузками конфигурации.
	GetChanges string
	//--config-dump-info-for-changes <путь> ‑ путь к файлу ConfigDumpInfo.xml, который используется для формирования файла изменений между двумя выгрузками конфигурации.
	ConfigDumpInfo string
	//--list-file <файл> ‑ выгрузить только объекты метаданных и/или внешние свойства, указанные в файле, вне зависимости от того были они изменены или нет.
	ListFile string
}

func (c DumpCfgToFiles) Command() string {
	return c.cmdConfig.Command() + " " + "dump-config-to-files"
}

func (c DumpCfgToFiles) Args() []string {

	v := []string{"--dir", c.Dir}
	if len(c.Extension) > 0 {
		v = append(v, "--extension", c.Extension)
	}
	if c.AllExtension {
		v = append(v, "--all-extensions")
	}
	if len(c.Format) > 0 {
		v = append(v, "--format", c.Format)
	}
	if c.Update {
		v = append(v, "--update")
	}
	if c.Force {
		v = append(v, "--force")
	}
	if len(c.GetChanges) > 0 {
		v = append(v, "--get-сhanges", c.GetChanges)
	}
	if len(c.ConfigDumpInfo) > 0 {
		v = append(v, "--config-dump-info-for-changes", c.ConfigDumpInfo)
	}
	if len(c.ListFile) > 0 {
		v = append(v, "--list-file", c.ListFile)
	}

	return v
}

//load-config-from-files
//Команда позволяет выполнить загрузку конфигурации из xml-файлов. Для команды доступны следующие параметры:
type LoadCfgFromFiles struct {
	cmdConfig

	//--dir <путь> ‑ содержит путь к каталогу, в который будет загружена конфигурация. Параметр является обязательным.
	Dir string `json:"dir"`
	//--extension <имя расширения> ‑ содержит имя расширения, которое будет выгружено в файл.
	Extension string `json:"extension"`
	//--all-extensions ‑ если параметр указан, то в файлы будут выгружены все расширения конфигурации.
	AllExtension bool `json:"all-extension"`
	//--format [hierarchical|plain] ‑ определяет формат выгрузки. По умолчанию используется иерархический формат выгрузки (hierarchical).
	Format string
	//--update-config-dump-info ‑ после окончания загрузки создать в директории файл ConfigDumpInfo.xml, соответствующий загруженной конфигурации.
	UpdateConfigDumpInfo bool
	//--files <файл[, файл]> ‑ список файлов, которые требуется загрузить.
	//	Файла разделяются запятыми.
	//	Пути к файлам указываются относительно каталога загрузки.
	//	Абсолютные пути не поддерживаются.
	//	При использовании параметра --list-file, данный параметр не используется.
	Files []string
	//--list-file <файл> ‑ путь к файлу, в котором перечислены загружаемые файлы.
	//Одна строка соответствует одному файлу.
	//Пути к файлам указываются относительно каталога загрузки.
	//Абсолютные пути не поддерживаются. При использовании параметра --files, данный параметр не используется.
	ListFile string
}

func (c LoadCfgFromFiles) Command() string {
	return c.cmdConfig.Command() + " " + "load-config-from-files"
}

func (c LoadCfgFromFiles) Args() []string {

	v := []string{"--dir", c.Dir}
	if len(c.Extension) > 0 {
		v = append(v, "--extension", c.Extension)
	}
	if c.AllExtension {
		v = append(v, "--all-extensions")
	}
	if len(c.Format) > 0 {
		v = append(v, "--format", c.Format)
	}
	if c.UpdateConfigDumpInfo {
		v = append(v, "--update-config-dump-info")
	}

	switch {
	case len(c.Files) > 0:
		v = append(v, "--files", strings.Join(c.Files, ", "))
	case len(c.ListFile) > 0:
		v = append(v, "--list-file", c.ListFile)
	}

	return v
}

//dump-cfg
//Команда позволяет выполнить выгрузку конфигурации или расширения в файл.
type DumpCfg struct {
	cmdConfig
	//--file <путь> ‑ путь к файлу конфигурации (cf-файл) или расширению (cfe-файл).
	File string `json:"file"`
	//--extension <имя расширения> ‑ содержит имя расширения, которое будет выгружено в файл.
	Extension string `json:"extension"`
}

func (c DumpCfg) Command() string {
	return c.cmdConfig.Command() + " " + "dump-cfg"
}

func (c DumpCfg) Args() []string {

	v := []string{"--file", c.File}
	if len(c.Extension) > 0 {
		v = append(v, "--extension", c.Extension)
	}
	return v
}

//load-cfg
//Команда позволяет выполни ть загрузку конфигурации или расширения из файла.
type LoadCfg struct {
	cmdConfig
	//--file <путь> ‑ путь к файлу конфигурации (cf-файл) или расширению (cfe-файл).
	File string `json:"file"`
	//--extension <имя расширения> ‑ содержит имя расширения, которое будет загружено из файла.
	Extension string `json:"extension"`
}

func (c LoadCfg) Command() string {
	return c.cmdConfig.Command() + " " + "dump-cfg"
}

func (c LoadCfg) Args() []string {

	v := []string{"--file", c.File}
	if len(c.Extension) > 0 {
		v = append(v, "--extension", c.Extension)
	}
	return v
}

//dump-external-data-processor-or-report-to-files
//Команда позволяет выполнить выгрузку внешних обработок или отчетов в xml-файлы.
type DumpToFilesEpfOrErf struct {
	cmdConfig
	//--file <путь> ‑ одержит имя файла, который будет выступать в роли корневого файла выгрузки внешней обработки/отчета в формате XML.
	//Параметр является обязательным.
	File string `json:"file"`
	//--ext-file <файл> ‑ полное имя файла с выгружаемой внешней обработкой (*.epf) или отчетом (*.erf).
	Ext string `json:"ext-file"`
	//--format [hierarchical|plain] ‑ определяет формат выгрузки. По умолчанию используется иерархический формат выгрузки (hierarchical).
	Format string `json:"format"`
}

func (c DumpToFilesEpfOrErf) Command() string {
	return c.cmdConfig.Command() + " " + "dump-external-data-processor-or-report-to-files"
}

func (c DumpToFilesEpfOrErf) Args() []string {

	v := []string{"--file", c.File}
	if len(c.Ext) > 0 {
		v = append(v, "--ext-file", c.Ext)
	}
	if len(c.Format) > 0 {
		v = append(v, "--format", c.Format)
	}
	return v
}

//load-external-data-processor-or-report-from-files
//Команда позволяет выполнить загрузку внешних обработок или отчетов из xml-файлов.
type LoadFromFilesEpfOrErf struct {
	cmdConfig
	//--file <путь> ‑ одержит имя файла, который будет выступать в роли корневого файла выгрузки внешней обработки/отчета в формате XML.
	//Параметр является обязательным.
	File string `json:"file"`
	//--ext-file <файл> ‑ полное имя файла с загружаемой внешней обработкой (*.epf) или отчетом (*.erf).
	Ext string `json:"ext-file"`
}

func (c LoadFromFilesEpfOrErf) Command() string {
	return c.cmdConfig.Command() + " " + "load-external-data-processor-or-report-from-files"
}

func (c LoadFromFilesEpfOrErf) Args() []string {

	v := []string{"--file", c.File}
	if len(c.Ext) > 0 {
		v = append(v, "--ext-file", c.Ext)
	}
	return v
}

//update-db-cfg
//Команда позволяет выполнить обновление конфигурации базы данных.
//При выполнении команды update-db-cfg используется следующий алгоритм работы:
//В том случае, если невозможно монопольно заблокировать базу данных и
//динамическое обновление возможно (не указан параметр --dynamic-disable) ‑
//будет выполнено динамическое обновление конфигурации базы данных.
//В том случае, если монопольно заблокировать базу данных возможно,
//но не требуется выполнять реструктуризацию базы данных ‑
//будет выполнено обычное обновление конфигурации базы данных.
//В том случае, если монопольно заблокировать базу данных возможно и
//требуется выполнить реструктуризацию базы данных, то будет выполнена следующая последовательность действий:
//  выполняется определение списка измененных объектов;
//  полученный список отображается в консоли;
//если указана необходимость подтверждения пользователя на принятие изменений (параметр --prompt-confirmation) ‑ пользователю задается соответствующий вопрос;
//если ответ на вопрос отрицательный ‑ обновление отменяется с выдачей соответствующего уведомления;
//если ответ утвердительный или запрос не требовался ‑ обновление продолжается штатным образом.
type UpdateDbCfg struct {
	cmdConfig
	//--prompt-confirmation ‑ определяет необходимость запроса у пользователя подтверждения о принятии изменений при реструктуризации информационной базы.
	PromptConfirmation bool
	//--dynamic-enable ‑ сначала выполняется попытка динамического обновления, если она завершена неудачно, будет запущено фоновое обновление.
	DynamicEnable bool
	//--dynamic-disable ‑ указание данного параметра запрещает динамическое обновление.
	DynamicDisable bool
	//--warnings-as-errors ‑ при указании данного параметра все предупреждения,
	//которые могут возникнуть при обновлении конфигурации базы данных, будут считаться ошибками.
	WarningsAsErrors bool
	//--background-start ‑ при указании данного параметра будет запущено фоновое обновление конфигурации, а текущий сеанс будет завершен.
	BackgroundStart bool
	//--background-cancel ‑ при указании данного параметра происходит отмена запущенного фонового обновления конфигурации базы данных.
	BackgroundCancel bool
	//--background-finish ‑ при указании данного параметра запущенное фоновое обновление конфигурации базы данных будет завершено.
	BackgroundFinish bool
	//При этом на базу данных будет наложена монопольная блокировка и проведена финальная фаза обновления.
	//--background-resume ‑ при указании данного параметра система продолжает фоновое обновление конфигурации базы данных, приостановленное ранее.
	BackgroundResume bool
	//--server ‑ данный параметр указывает, что обновление конфигурации базы данных необходимо выполнить на стороне сервер «1С:Предприятия».
	Server bool
	//--extension <имя расширения> ‑ имя расширения.
	Extension string
}

func (c UpdateDbCfg) Command() string {
	return c.cmdConfig.Command() + " " + "update-db-cfg"
}

func (c UpdateDbCfg) Args() []string {

	var v []string

	if c.PromptConfirmation {
		v = append(v, "--prompt-confirmation")
	}
	if c.DynamicEnable {
		v = append(v, "--dynamic-enable")
	}
	if c.DynamicDisable {
		v = append(v, "--dynamic-disable")
	}
	if c.WarningsAsErrors {
		v = append(v, "--warnings-as-errors")
	}
	if c.BackgroundStart {
		v = append(v, "--background-start")
	}
	if c.BackgroundCancel {
		v = append(v, "--background-cancel")
	}
	if c.BackgroundFinish {
		v = append(v, "--background-finish")
	}
	if c.BackgroundResume {
		v = append(v, "--background-resume")
	}

	if len(c.Extension) > 0 {
		v = append(v, "--extension", c.Extension)
	}
	return v
}

//manage-cfg-support
//Команда позволяет снимать конфигурация с поддержки
type ManageCfgSupport struct {
	cmdConfig
	//--disable-support ‑ указывает на необходимость снятия конфигурации с поддержки.
	//При отсутствии параметра генерируется ошибка.
	DisableSupport bool
	//--force ‑ выполнить снятие конфигурации с поддержки даже в том случае, если в конфигурации запрещены изменения.
	//При отсутствии параметра будет сгенерирована ошибка, если попытка снятия с поддержки будет выполняться для конфигурации,
	//для которой в интерактивном режиме управления поддержкой запрещены изменения.
	Force bool
}

func (c ManageCfgSupport) Command() string {
	return c.cmdConfig.Command() + " " + "manage-cfg-support"
}

func (c ManageCfgSupport) Args() []string {

	var v []string

	if c.DisableSupport {
		v = append(v, "--bdisable-support")
	}
	if c.Force {
		v = append(v, "--force")
	}
	return v
}

type cmdConfigExtension struct{}

func (c cmdConfigExtension) Command() string {
	return fmt.Sprintf("%s %s", "config", "extensions")
}

type ExtensionPurposeType string

const (
	ExtensionPurposeCustomization ExtensionPurposeType = "customization"
	ExtensionPurposeAddOn                              = "add-on"
	ExtensionPurposePatch                              = "patch"
)

//create
//Команда предназначена для создания расширения в информационной базе.
//Расширение создается пустым.
//Для загрузки расширения следует использовать команду config load-cfg
//или config load-config-from-files.
type CreateExtension struct {
	cmdConfigExtension
	//--extension <имя> ‑ задает имя расширения. Параметр является обязательным.
	Name string
	//--name-prefix <префикс> ‑ задает префикс имени для расширения. Параметр является обязательным.
	Prefix string
	//--synonym <синоним> ‑ синоним имени расширения. Многоязычная строка в формате функции Nstr().
	Synonym string
	//--purpose <назначение> ‑ назначение расширения. <Назначение> может принимать следующие значения:
	// customization ‑ назначение Адаптация (значение по умолчанию);
	// add-on ‑ назначение Дополнение;
	// patch ‑ назначение Исправление.
	Purpose ExtensionPurposeType
}

func (c CreateExtension) Command() string {
	return fmt.Sprintf("%s %s", c.cmdConfigExtension.Command(), "create")
}

func (c CreateExtension) Args() []string {

	var v []string
	v = append(v, "--extension", c.Name)
	v = append(v, "--name-prefix", c.Prefix)

	if len(c.Synonym) > 0 {
		v = append(v, "--synonym", string(c.Synonym))
	}

	if len(c.Purpose) > 0 {
		v = append(v, "--purpose", string(c.Purpose))
	}
	return v
}

//create
//Команда предназначена для создания расширения в информационной базе.
//Расширение создается пустым.
//Для загрузки расширения следует использовать команду config load-cfg
//или config load-config-from-files.
type DeleteExtension struct {
	cmdConfigExtension
	// --extension <имя> ‑ задает имя удаляемого расширения.
	Name string
	//--all-extensions ‑ указывает, что необходимо удалить все расширения.
	All bool
}

func (c DeleteExtension) Command() string {
	return fmt.Sprintf("%s %s", c.cmdConfigExtension.Command(), "delete")
}

func (c DeleteExtension) Args() []string {

	var v []string
	v = append(v, "--extension", c.Name)

	if c.All {
		v = append(v, "--all-extensions")
	}

	return v
}

//Группа команд properties
//Назначение группы команд
//Группа команд config extensions properties позволяет задавать и получать свойства расширения.
//
//get
//Команда предназначена для получения свойств расширения, расположенного в информационной базе.
type GetExtensionProperties struct {
	cmdConfigExtension
	//--extension <имя> ‑ задает имя расширения, для которого необходимо получить свойства.
	Extension string
	//--all-extensions ‑ указывает, что необходимо получить свойства всех расширений, загруженных в информационную базу.
	All bool
}

func (c GetExtensionProperties) Command() string {
	return fmt.Sprintf("%s %s %s", c.cmdConfigExtension.Command(), "properties", "get")
}

func (c GetExtensionProperties) Args() []string {

	var v []string
	v = append(v, "--extension", c.Extension)

	if c.All {
		v = append(v, "--all-extensions")
	}

	return v
}

type ExtensionPropertiesScopeType string

const (
	ExtensionPropertiesScopeInfobase       ExtensionPropertiesScopeType = "infobase"
	ExtensionPropertiesScopeDataSeparation                              = "data-separation"
)

func (t ExtensionPropertiesScopeType) String() string {
	return string(t)
}

type ExtensionPropertiesBoolType string

func (t ExtensionPropertiesBoolType) String() string {
	return string(t)
}

const (
	ExtensionPropertiesBoolYes    ExtensionPropertiesBoolType = "yes"
	ExtensionPropertiesBoolNo     ExtensionPropertiesBoolType = "no"
	ExtensionPropertiesBoolNotSet                             = ""
)

//set
//Команда предназначена для установки свойств расширения, расположенного в информационной базе.
type SetExtensionProperties struct {
	cmdConfigExtension
	//--extension <имя> ‑ задает имя расширения, для которого необходимо получить свойства.
	Extension string
	//--active <режим> ‑ определяет активность расширения. <Режим> может принимать следующие значения:
	// yes ‑ расширение активно.
	// no ‑ расширение не активно.
	Active ExtensionPropertiesBoolType
	//--safe-mode <режим> ‑ определяет работу в безопасном режиме. <Режим> может принимать следующие значения:
	// yes ‑ расширение работает в безопасном режиме.
	// no ‑ расширение работает в небезопасном режиме. В этом случае имя профиля безопасности автоматически сбрасывается (имя профиля устанавливается равным пустой строке).
	SafeMode ExtensionPropertiesBoolType
	//--security-profile-name <профиль> ‑ определяет имя профиля безопасности, под управлением которого работает расширение.
	//Если задается имя профиля безопасности, то автоматически устанавливается и признак работы в безопасном режиме.
	SecurityProfileName string
	//--unsafe-action-protection <режим> ‑ определяет режим защиты от опасных действий. <Режим> может принимать следующие значения:
	// yes ‑ защита от опасных действий в расширении включена.
	// no ‑ защита от опасных действий в расширении отключена.
	UnsafeActionProtection ExtensionPropertiesBoolType
	//--used-in-distributed-infobase <режим> ‑ определяет возможность работы расширения в распределенной информационной базе. <Режим> может принимать следующие значения:
	// yes ‑ расширение используется в распределенной информационной базе.
	// no ‑ расширение не используется в распределенной информационной базе.
	UsedInDistributedInfobase ExtensionPropertiesBoolType
	//--scope <область действия> ‑ область действия расширения. <Область действия> может принимать следующие значения:
	// infobase ‑ расширение действительно для все информационной базы.
	// data-separation ‑ расширение действительно для области данных.
	Scope ExtensionPropertiesScopeType
}

func (c SetExtensionProperties) Command() string {
	return fmt.Sprintf("%s %s %s", c.cmdConfigExtension.Command(), "properties", "get")
}

func (c SetExtensionProperties) Args() []string {

	var v []string
	v = append(v, "--extension", c.Extension)

	if c.Active != ExtensionPropertiesBoolNotSet {
		v = append(v, "--active", c.Active.String())
	}

	if c.SafeMode != ExtensionPropertiesBoolNotSet {
		v = append(v, "--safe-mode", c.SafeMode.String())
	}

	if c.UnsafeActionProtection != ExtensionPropertiesBoolNotSet {
		v = append(v, "--unsafe-action-protection", c.UnsafeActionProtection.String())
	}

	if c.UsedInDistributedInfobase != ExtensionPropertiesBoolNotSet {
		v = append(v, "--used-in-distributed-infobase", c.UsedInDistributedInfobase.String())
	}

	if len(c.Scope) > 0 {
		v = append(v, "--scope", c.Scope.String())

	}

	return v
}

//Команды группы infobase-tools отвечают за получение сервисной информации об информационной базе.
type cmdInfobaseTools struct{}

func (c cmdInfobaseTools) Command() string {
	return "infobase-tools"
}

func (c cmdInfobaseTools) Args() []string {
	var v []string
	return v
}

//debug-info
//C помощью данной команды получить информацию о настройках отладчика для информационной базы.
type DebigInfo struct {
	cmdInfobaseTools
}

func (c DebigInfo) Command() string {
	return fmt.Sprintf("%s %s", c.cmdInfobaseTools.Command(), "debug-info")
}

//data-separation-common-attribute-list
//Данная команда позволяет получить список имен разделителей информационной базы.
type DataSeparationList struct {
	cmdInfobaseTools
}

func (c DataSeparationList) Command() string {
	return fmt.Sprintf("%s %s", c.cmdInfobaseTools.Command(), "data-separation-common-attribute-list")
}

//dump-ib
//Команда предназначена для выполнения выгрузки информационной базы в dt-файл. Допустимо использование следующих параметров:
type DumpInfobase struct {
	cmdInfobaseTools

	// --file <имя файла> ‑ определяет имя dt-файла.
	File string
}

func (c DumpInfobase) Command() string {
	return fmt.Sprintf("%s %s", c.cmdInfobaseTools.Command(), "dump-ib")
}

func (c DumpInfobase) Args() []string {
	return []string{"--file", c.File}
}

//restore-ib
//Команда предназначена для выполнения загрузки информационной базы из dt-файл. Допустимо использование следующих параметров:
type RestoreInfobase struct {
	cmdInfobaseTools

	// --file <имя файла> ‑ определяет имя dt-файла.
	File string
}

func (c RestoreInfobase) Command() string {
	return fmt.Sprintf("%s %s", c.cmdInfobaseTools.Command(), "restore-ib")
}

func (c RestoreInfobase) Args() []string {
	return []string{"--file", c.File}
}

//erase-data
//Команда выполняет удаление данных информационной базы.
type EraseInfobaseDate struct {
	cmdInfobaseTools
}

func (c EraseInfobaseDate) Command() string {
	return fmt.Sprintf("%s %s", c.cmdInfobaseTools.Command(), "erase-data")
}
