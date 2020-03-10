package v8run

type RepositoryRightType string
type RepositorySupportEditObjectsType string

const (
	STORAGE_RIGHT_READ            RepositoryRightType = "ReadOnly"
	STORAGE_RIGHT_LOCK                                = "LockObjects"
	STORAGE_RIGHT_MANAGE_VERSIONS                     = "ManageConfigurationVersions"
	STORAGE_RIGHT_ADMIN                               = "Administration"
)

const (
	REPOSITORY_SUPPORT_NOT_EDITABLE  RepositorySupportEditObjectsType = "ObjectNotEditable"
	REPOSITORY_SUPPORT_IS_EDITABLE                                    = "ObjectIsEditableSupportEnabled"
	REPOSITORY_SUPPORT_NOT_SUPPORTED                                  = "ObjectNotSupported"
)

func (t RepositorySupportEditObjectsType) MarshalV8() (string, error) {
	return string(t), nil
}

func (t RepositoryRightType) MarshalV8() (string, error) {
	return string(t), nil
}

type Repository struct {
	///ConfigurationRepositoryF <каталог хранилища>
	//— указание имени каталога хранилища.
	Path string `v8:"/ConfigurationRepositoryF" json:"path"`

	///ConfigurationRepositoryN <имя>
	//— указание имени пользователя хранилища.
	User string `v8:"/ConfigurationRepositoryN" json:"user"`

	///ConfigurationRepositoryP <пароль>
	//— указание пароля пользователя хранилища.
	Password string `v8:"/ConfigurationRepositoryP, optional" json:"password"`
}

///ConfigurationRepositoryAddUser [-Extension <имя расширения>] -User <Имя> -Pwd <Пароль> -Rights <Права> [-RestoreDeletedUser]
//— создание пользователя хранилища конфигурации.
// Пользователь, от имени которого выполняется подключение к хранилищу,
// должен обладать административными правами.
// Если пользователь с указанным именем существует, то пользователь добавлен не будет.
type RepositoryAddUserOptions struct {
	*Designer   `v8:",inherit" json:"designer"`
	*Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryAddUser" json:"-"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//-User — Имя создаваемого пользователя.
	NewUser string `v8:"-User" json:"user"`

	//-Pwd — Пароль создаваемого пользователя.
	NewPassword string `v8:"-pwd, optional" json:"pwd"`

	//-Rights — Права пользователя. Возможные значения:
	//	ReadOnly — право на просмотр,
	//	LockObjects — право на захват объектов,
	//	ManageConfigurationVersions — право на изменение состава версий,
	//	Administration — право на административные функции.
	//
	Rights RepositoryRightType `v8:"-Rights" json:"rights"`

	//-RestoreDeletedUser — Если обнаружен удаленный пользователь с таким же именем, он будет восстановлен.
	RestoreDeletedUser bool `v8:"-RestoreDeletedUser" json:"restore_deleted_user"`
}

//ConfigurationRepositoryCreate
///ConfigurationRepositoryCreate [-Extension <имя расширения>] [-AllowConfigurationChanges
//-ChangesAllowedRule <Правило поддержки> -ChangesNotRecommendedRule <Правило поддержки>] [-NoBind]
//— предназначен для создания хранилища конфигурации. Доступны следующие параметры:
//Пример:
//
//DESIGNER /F "D:\V8\Cfgs83\ИБ83" /ConfigurationRepositoryF "D:\V8\Cfgs83" /ConfigurationRepositoryN "Администратор"
// /ConfigurationRepositoryP "123456" /ConfigurationRepositoryCreate - AllowConfigurationChanges
// -ChangesAllowedRule ObjectNotEditable -ChangesNotRecommendedRule ObjectNotEditable
type RepositoryCreateOptions struct {
	*Designer   `v8:",inherit" json:"designer"`
	*Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryCreate" json:"-"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//-AllowConfigurationChanges — если конфигурация находится на поддержке без возможности изменения, будет включена возможность изменения.
	AllowConfigurationChanges bool `v8:"-AllowConfigurationChanges" json:"allow_configuration_changes"`

	//-ChangesAllowedRule <Правило поддержки> — устанавливает правило поддержки для объектов,
	// для которых изменения разрешены поставщиком. Может быть установлено одно из следующих правил:
	//	ObjectNotEditable - Объект поставщика не редактируется,
	//	ObjectIsEditableSupportEnabled - Объект поставщика редактируется с сохранением поддержки,
	//	ObjectNotSupported - Объект поставщика снят с поддержки.
	ChangesAllowedRule RepositorySupportEditObjectsType `v8:"-ChangesAllowedRule, optional" json:"changes_allowed_rule"`

	//-ChangesNotRecommendedRule — устанавливает правило поддержки для объектов,
	// для которых изменения не рекомендуются поставщиком. Может быть установлено одно из следующих правил:
	//	ObjectNotEditable - Объект поставщика не редактируется,
	//	ObjectIsEditableSupportEnabled - Объект поставщика редактируется с сохранением поддержки,
	//	ObjectNotSupported - Объект поставщика снят с поддержки.
	ChangesNotRecommendedRule RepositorySupportEditObjectsType `v8:"-ChangesAllowedRule, optional" json:"changes_not_recommended_rule"`

	//-NoBind — К созданному хранилищу подключение выполнено не будет.
	NoBind bool `v8:"-NoBind" json:"no_bind"`
}
