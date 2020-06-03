package designer

import (
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
)

var _ types.Command = (*Designer)(nil)

type Designer struct {
	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
	Visible                bool `v8:"/Visible" json:"visible"`
}

func (d Designer) Command() string {
	return types.COMMAND_DESIGNER
}

func (d Designer) Check() error {

	return nil
}

func (d Designer) Values() *types.Values {
	v, _ := marshaler.Marshal(d)
	return v

}

func NewDesigner() Designer {

	d := Designer{
		DisableStartupDialogs:  true,
		DisableStartupMessages: true,
		Visible:                false,
	}

	return d
}

type UpdateCfgOptions struct {
	Designer `v8:",inherit" json:"designer"`
	//<имя cf | cfu-файла>
	File string `v8:"/UpdateCfg" json:"file"`

	// <имя файла настроек> — содержит имя файла настроек объединения.
	Settings string `v8:"-Settings" json:"settings"`

	// если в настройках есть объекты, не включенные в список обновляемых и отсутствующие в основной конфигурации,
	// на которые есть ссылки из объектов, включенных в список, то такие объекты также помечаются для обновления,
	// и выполняется попытка продолжить обновление.
	IncludeObjectsByUnresolvedRefs bool `v8:"-IncludeObjectsByUnresolvedRefs" json:"include_objects_by_unresolved_refs"`

	//— очищение ссылок на объекты, не включенные в список обновляемых.
	ClearUnresolvedRefs bool `v8:"-ClearUnresolvedRefs" json:"clear_unresolved_refs"`

	//— Если параметр используется, обновление будет выполнено несмотря на наличие предупреждений:
	//о применении настроек,
	//о дважды измененных свойствах, для которых не был выбран режим объединения,
	//об удаляемых объектах, на которые найдены ссылки в объектах, не участвующие в объединении.
	//Если параметр не используется, то в описанных случаях объединение будет прервано.
	Force bool `v8:"-Force" json:"force"`

	//— вывести список всех дважды измененных свойств.
	DumpListOfTwiceChangedProperties bool `v8:"-DumpListOfTwiceChangedProperties" json:"dump_list_of_twice_changed_properties"`

	UpdateDBCfg *UpdateDBCfgOptions `v8:",inherit" json:"update_db"`
}

func (d UpdateCfgOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v

}

func (d UpdateCfgOptions) WithUpdateDBCfg(upd UpdateDBCfgOptions) UpdateCfgOptions {

	UpdateDBCfg := &upd

	return UpdateCfgOptions{
		Designer:                         d.Designer,
		File:                             d.File,
		Settings:                         d.Settings,
		IncludeObjectsByUnresolvedRefs:   d.IncludeObjectsByUnresolvedRefs,
		ClearUnresolvedRefs:              d.ClearUnresolvedRefs,
		Force:                            d.Force,
		DumpListOfTwiceChangedProperties: d.DumpListOfTwiceChangedProperties,
		UpdateDBCfg:                      UpdateDBCfg,
	}

}

///UpdateDBCfg [–Dynamic<Режим>] [-BackgroundStart] [-BackgroundCancel]
//[-BackgroundFinish [-Visible]] [-BackgroundSuspend] [-BackgroundResume]
//[-WarningsAsErrors] [-Server [-v1|-v2]][-Extension <имя расширения>]
type UpdateDBCfgOptions struct {
	Designer `v8:",inherit" json:"designer"`

	command struct{} `v8:"/UpdateDBCfg" json:"update_db_configuration"`

	//-Dynamic<Режим> — признак использования динамического обновления. Режим может принимать следующие значения
	//-Dynamic+ — Значение параметра по умолчанию.
	// Сначала выполняется попытка динамического обновления, если она завершена неудачно, будет запущено фоновое обновление.
	//-Dynamic–  — Динамическое обновление запрещено.
	Dynamic bool `v8:"-Dynamic, no_span, bool_false=-, bool_true=+" json:"dynamic"`

	//-BackgroundStart [-Dynamic<Режим>] — будет запущено фоновое обновление конфигурации,
	// текущий сеанс будет завершен. Если обновление уже выполняется, будет выдана ошибка.
	//-Dynamic+ — Значение параметра по умолчанию.
	// Сначала выполняется попытка динамического обновления, если она завершена неудачно,
	// будет запущено фоновое обновление.
	//-Dynamic–  — Динамическое обновление запрещено.
	BackgroundStart bool `v8:"-BackgroundCancel" json:"background_start"`

	//-BackgroundCancel — отменяет запущенное фоновое обновление конфигурации базы данных.
	// Если фоновое обновление не запущено, будет выдана ошибка.
	BackgroundCancel bool `v8:"-BackgroundCancel"  json:"background_cancel"`

	//-BackgroundFinish — запущенное фоновое обновление конфигурации базы данных будет завершено:
	// при этом будет наложена монопольная блокировка и проведена финальная фаза обновления.
	// Если фоновое обновление конфигурации не запущено или переход к завершающей фазе обновления не возможен, будет выдана ошибка.
	// Возможно использование следующих параметров:
	//-Visible — На экран будет выведен диалоговое окно с кнопками Отмена, Повторить, Завершить сеансы и повторить.
	// В случае невозможности завершения фонового обновления, если данная опция не указана, выполнение обновления будет завершено с ошибкой..
	BackgroundFinish bool `v8:"-BackgroundFinish" json:"background_finish"`

	//-BackgroundResume — продолжает фоновое обновление конфигурации базы данных, приостановленное ранее.
	BackgroundResume bool `v8:"-BackgroundResume" json:"background_resume"`

	//-BackgroundSuspend — приостанавливает фоновое обновление конфигурации на паузу.
	// Если фоновое обновление не запущено, будет выдана ошибка.
	BackgroundSuspend bool `v8:"-BackgroundSuspend" json:"background_suspend"`

	//-WarningsAsErrors —  все предупредительные сообщения будут трактоваться как ошибки.
	WarningsAsErrors bool `v8:"-WarningsAsErrors" json:"warnings_as_errors"`

	//-Server — обновление будет выполняться на сервере (имеет смысл только на сервере).
	// Если параметр используется вместе с фоновым обновлением, то:
	//
	//Фаза актуализации всегда выполняется на сервере.
	//Фаза обработки и фаза принятия изменения могут выполняться как на клиенте, так и на сервере.
	//Допускается запуск фонового обновления на стороне клиента, а завершение - на стороне сервера, и наоборот.
	//Не используется 2-я версия механизма реструктуризации (игнорируется параметр -v2, если таковой указан).
	//Если не указана версия механизма реструктуризации (-v1 или -v2),
	// то будет использоваться механизм реструктуризации той версии, которая указана в файле conf.cfg.
	// В противном случае будет использована указанная версия механизма.
	// Если указана 2-я версия механизма реструктуризации, но использование этой версии конфликтует с другими параметрами
	// – будет использована 1-я версия.
	Server bool `v8:"-Server" json:"server"`

	//-Extension <Имя расширения> — будет выполнено обновление расширения с указанным именем.
	// Если расширение успешно обработано возвращает код возврата 0,
	// в противном случае (если расширение с указанным именем не существует или в процессе работы произошли ошибки) — 1.
	Extension string `v8:"-Extension, optional" json:"extension"`
}

func (d UpdateDBCfgOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v

}

func (d UpdateDBCfgOptions) WithExtension(extension string) UpdateDBCfgOptions {

	return UpdateDBCfgOptions{
		Designer:          d.Designer,
		Dynamic:           d.Dynamic,
		BackgroundStart:   d.BackgroundStart,
		BackgroundCancel:  d.BackgroundCancel,
		BackgroundFinish:  d.BackgroundFinish,
		BackgroundResume:  d.BackgroundResume,
		BackgroundSuspend: d.BackgroundSuspend,
		WarningsAsErrors:  d.WarningsAsErrors,
		Server:            d.Server,
		Extension:         extension,
	}

}
