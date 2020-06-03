package designer

import (
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
	"strings"
)

type LoadCfgOptions struct {
	Designer `v8:",inherit" json:"designer"`

	File string `v8:"/LoadCfg" json:"file"`

	Extension   string              `v8:"-Extension, optional" json:"extension"`
	UpdateDBCfg *UpdateDBCfgOptions `v8:",inherit" json:"update_db_cfg"`
}

func (d LoadCfgOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v

}

func (d LoadCfgOptions) WithUpdateDBCfg(upd UpdateDBCfgOptions) LoadCfgOptions {

	UpdateDBCfg := &upd

	return LoadCfgOptions{
		Designer:    d.Designer,
		File:        d.File,
		Extension:   d.Extension,
		UpdateDBCfg: UpdateDBCfg,
	}

}

func (d LoadCfgOptions) WithExtension(extension string) LoadCfgOptions {

	return LoadCfgOptions{
		Designer:    d.Designer,
		File:        d.File,
		Extension:   extension,
		UpdateDBCfg: d.UpdateDBCfg,
	}

}

type FileList []string

func (t FileList) MarshalV8() (string, error) {
	return strings.Join(t, ","), nil
}

//LoadConfigFromFiles <каталог загрузки> [-Extension <имя расширения>]
//[-AllExtensions][-files "<файлы>"][-listFile <файл списка>][-format <режим>] [-updateConfigDumpInfo]
//— загрузка конфигурации из файлов. Загрузка расширения в основную конфигурацию (и наоборот) не поддерживается.
type LoadConfigFromFiles struct {
	Designer `v8:",inherit" json:"designer"`

	//<каталог загрузки> — каталог, содержащий XML-файлы конфигурации;
	Dir string `v8:"/LoadConfigFromFiles" json:"dir"`

	//Extension <Имя расширения> — обработка расширения с указанным именем.
	//Если расширение успешно обработано возвращает код возврата 0,
	//в противном случае (если расширение с указанным именем не существует или в процессе работы произошли ошибки) — 1;
	//Важно! Если указанное расширение подключено к хранилищу,
	//возможна только частичная загрузка, если соответствующие объекты были захвачены в хранилище.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//AllExtensions — загрузка только расширений (всех).
	//Если требуемое расширение не существует, оно будет создано.
	//Для каждого подкаталога указанного каталога будет выполнена попытка создать расширение.
	//При попытке загрузить расширение в основную конфигурацию или наоборот, будет выведена ошибка.
	AllExtensions bool `v8:"-AllExtensions, optional" json:"all-extensions"`

	//updateConfigDumpInfo — указывает, что в конце загрузки
	//в каталоге будет создан файл версий ConfigDumpInfo.xml,
	//соответствующий загруженной конфигурации.
	//Если выполняется частичная загрузка (используется параметр -files или -listFile),
	//файл версий будет обновлен.
	UpdateDumpInfo bool `v8:"-updateConfigDumpInfo" json:"update_config_dump_info"`

	//files — содержит список файлов, которые требуется загрузить.
	//Список разделяется запятыми.
	//Не используется, если указан параметр -listFile.
	//При запуске в режиме агента путь к загружаемым файлам должен быть относительным.
	Files FileList `v8:"-files" json:"files"`

	//listFile — указывает файл, в котором перечислены файлы, которые требуется загрузить. Не используется, если указан параметр -files. При запуске в режиме агента путь к загружаемым файлам должен быть относительным.
	//Указываемый файл должен удовлетворять следующим требованиям:
	//- Файл должен быть в кодировке UTF-8.
	//- Имена файлов должны быть указаны через перенос (поддерживаются символы переноса \r\n ("следующая строка") и \r ("возврат каретки")).
	//- Файл не должен содержать пустые строки между именами файлов.
	ListFile string `v8:"-listFile, optional" json:"list_file"`

	UpdateDBCfg *UpdateDBCfgOptions `v8:",inherit" json:"update_db_cfg"`
}

func (o LoadConfigFromFiles) Values() *types.Values {

	v, _ := marshaler.Marshal(o)
	return v
}

func (o LoadConfigFromFiles) WithExtension(extension string) LoadConfigFromFiles {

	newO := o
	newO.Extension = extension
	return newO

}

func (o LoadConfigFromFiles) WithFiles(files ...string) LoadConfigFromFiles {

	newO := o
	newO.Files = files
	return newO

}

func (o LoadConfigFromFiles) WithListFile(file string) LoadConfigFromFiles {

	newO := o
	newO.ListFile = file
	return newO

}

func (o LoadConfigFromFiles) WithUpdateDBCfg(upd UpdateDBCfgOptions) LoadConfigFromFiles {

	UpdateDBCfg := &upd
	newO := o

	newO.UpdateDBCfg = UpdateDBCfg

	return newO

}

func (o LoadConfigFromFiles) WithUpdateDumpInfo() LoadConfigFromFiles {

	newO := o

	newO.UpdateDumpInfo = true

	return newO

}

func (o LoadConfigFromFiles) WithAllExtension() LoadConfigFromFiles {

	newO := o
	newO.AllExtensions = true
	return newO

}

///LoadExternalDataProcessorOrReportFromFiles <путь к корневому файлу выгрузки> <путь к файлу внешней обработки или отчета>
//— загрузка внешних обработок или отчетов из файлов. Все параметры являются обязательными:
type LoadExternalDataFileFromFilesOptions struct {
	Designer `v8:",inherit" json:"designer"`

	command struct{} `v8:"/LoadExternalDataProcessorOrReportFromFiles" json:"-"`

	//<путь к корневому файлу выгрузки> — содержит путь к корневому файлу выгрузки внешний обработки или отчета в формате XML.
	Dir string `v8:",arg" json:"dir"`

	//<путь к файлу внешней обработки или отчета> — содержит путь к файлу внешней обработки или отчета,
	//в который будет записан результат загрузки из XML-файла.
	//Расширение результирующего файла всегда соответствует содержимому исходной выгрузки:
	//".epf" — для внешних обработок,
	//".erf" — для отчетов.
	//Если в качестве параметра задан файл с другим расширением,
	//то оно будет заменено на соответствующее.
	File string `v8:",arg" json:"file"`
}

func (o LoadExternalDataFileFromFilesOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(o)
	return v
}
