package designer

import (
	"github.com/Khorevaa/go-v8platform/marshaler"
	"github.com/Khorevaa/go-v8platform/types"
)

type DumpCfgOptions struct {
	Designer  `v8:",inherit" json:"designer"`
	File      string `v8:"/DumpCfg" json:"file"`
	Extension string `v8:"-Extension, optional" json:"extension"`
}

func (d DumpCfgOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v
}

func (d DumpCfgOptions) WithExtension(extension string) DumpCfgOptions {

	return DumpCfgOptions{
		Designer:  d.Designer,
		File:      d.File,
		Extension: extension,
	}

}

///DumpConfigToFiles <каталог выгрузки> [-Extension <имя расширения>]
//[-AllExtensions] [-format] [-update][-force][-getChanges <имя файла>]
//[-configDumpInfoForChanges <имя файла>][-configDumpInfoOnly]
//[-listFile <имя файла>]
//— выгрузка конфигурации в XML-файлы. При выгрузке будет создан файл версий (ConfigDumpInfo.xml).
type DumpConfigToFilesOptions struct {
	Designer `v8:",inherit" json:"designer"`

	//<каталог выгрузки> — каталог, в который будет выгружена конфигурация.
	Dir string `v8:"/DumpConfigToFiles" json:"dir"`

	//Extension <Имя расширения> — выгрузка расширения с указанным именем.
	//Если расширение успешно обработано возвращает код возврата 0,
	//в противном случае (если расширение с указанным именем не существует или в процессе работы произошли ошибки) — 1.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//AllExtensions — выгрузка только расширений (всех).
	//Для каждого расширения будет создан каталог, имя которого совпадает с именем расширения.
	AllExtensions bool `v8:"-AllExtensions, optional" json:"all-extensions"`

	//-force — Если текущая версия формата выгрузки не совпадает
	//с версией формата в файле версий, будет выполнена полная выгрузка.
	Force bool `v8:"-force" json:"force"`

	//update — указывает, что выгрузка будет обновлена (будут выгружены только файлы, версии которых отличаются от ранее выгруженных).
	//Файл версий (ConfigDumpInfo.xml) будет получен из текущего каталога выгрузки. Если текущая версия формата выгрузки не совпадает с версией формата в файле версий или если файл версий не найден, будет сгенерирована ошибка. По завершении выгрузки файл версий обновляется.
	//Возможно совместное использование с параметрами:
	Update bool `v8:"-update" json:"update"`

	//configDumpInfoForChanges <имя файла версий> — указывает файл версий,
	//который будет использован для сравнения изменений.
	//Имя файла версий должно быть указано.
	//Примечание. Данная опция используется только совместно с параметрами -update и -getChanges.
	ConfigDumpInfoForChanges string `v8:"-configDumpInfoForChanges, optional" json:"config_dump_info_for_changes"`
}

func (o DumpConfigToFilesOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(o)
	return v
}

func (o DumpConfigToFilesOptions) WithExtension(extension string) DumpConfigToFilesOptions {

	newO := o
	newO.Extension = extension
	return newO

}

func (o DumpConfigToFilesOptions) WithUpdate(force bool, configDumpInfo string) DumpConfigToFilesOptions {

	newO := o

	newO.Force = force
	newO.Update = true
	newO.ConfigDumpInfoForChanges = configDumpInfo

	return newO

}

func (o DumpConfigToFilesOptions) WithAllExtension() DumpConfigToFilesOptions {

	newO := o
	newO.AllExtensions = true
	return newO

}

///DumpConfigToFiles <каталог выгрузки> [-Extension <имя расширения>]
//[-AllExtensions] [-format] [-update][-force][-getChanges <имя файла>]
//[-configDumpInfoForChanges <имя файла>][-configDumpInfoOnly]
//[-listFile <имя файла>]
//— выгрузка конфигурации в XML-файлы. При выгрузке будет создан файл версий (ConfigDumpInfo.xml).
type GetChangesForConfigDumpOptions struct {
	Designer `v8:",inherit" json:"designer"`

	//<каталог выгрузки> — каталог, в который будет выгружена конфигурация.
	Dir string `v8:"/DumpConfigToFiles" json:"dir"`

	//Extension <Имя расширения> — выгрузка расширения с указанным именем.
	//Если расширение успешно обработано возвращает код возврата 0,
	//в противном случае (если расширение с указанным именем не существует или в процессе работы произошли ошибки) — 1.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//-force — Если текущая версия формата выгрузки не совпадает
	//с версией формата в файле версий, будет выполнена полная выгрузка.
	Force bool `v8:"-force" json:"force"`

	//configDumpInfoForChanges <имя файла версий> — указывает файл версий,
	//который будет использован для сравнения изменений.
	//Имя файла версий должно быть указано.
	//Примечание. Данная опция используется только совместно с параметрами -update и -getChanges.
	ConfigDumpInfoForChanges string `v8:"-configDumpInfoForChanges, optional" json:"config_dump_info_for_changes"`

	//getChanges <имя файла>  — в указанный файл будут выведены изменения
	//текущей конфигурации относительно выгрузки,
	//каталог которой указан перед данным параметром.
	//Изменения вычисляются относительно файла версий в текущем каталоге выгрузки. Имя файла должно быть указано.
	//Примечание. Может использоваться совместно с параметром configDumpInfoForChanges
	//- изменения будут вычислены относительно переданного файла версий.
	// Если при использовании параметром configDumpInfoForChanges файл версии не найден, будет сгенерирована ошибка.
	GetChanges string `v8:"-getChanges, optional" json:"get_changes"`
}

func (o GetChangesForConfigDumpOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(o)
	return v
}

func (o GetChangesForConfigDumpOptions) WithExtension(extension string) GetChangesForConfigDumpOptions {

	newO := o
	newO.Extension = extension
	return newO

}

func (o GetChangesForConfigDumpOptions) WithConfigDumpInfo(configDumpInfo string) GetChangesForConfigDumpOptions {

	newO := o
	newO.ConfigDumpInfoForChanges = configDumpInfo
	return newO

}

///DumpExternalDataProcessorOrReportToFiles <путь к корневому файлу выгрузки>
//<путь к файлу внешней обработки или отчета> [-Format Plain|Hierarchical]
//— выгрузка внешних обработок или отчетов в файлы. Доступны следующие параметры:
type DumpExternalDataFileToFilesOptions struct {
	Designer `v8:",inherit" json:"designer"`

	command struct{} `v8:"/DumpExternalDataProcessorOrReportToFiles" json:"-"`
	//<путь к корневому файлу выгрузки> —  содержит путь к корневому файлу выгрузки,
	//в который будут сохранены файлы в формате XML внешней обработки или отчета.
	Dir string `v8:",arg" json:"dir"`

	//<путь к файлу внешней обработки или отчета> — содержит путь к файлу внешней обработки (.epf) или отчета (.erf).
	File string `v8:",arg" json:"file"`
}

func (o DumpExternalDataFileToFilesOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(o)
	return v
}
