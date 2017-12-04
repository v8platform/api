package v8platform

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"../v8config-file"
	"../v8tools"
	"github.com/mcuadros/go-version"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const ВерсияПоУмолчанию = "8.3"
const этоWindows = runtime.GOOS == "windows"

type ВерсияПлатформы struct {
	Версия string
	Rac    string
	V8     string
}

func новаяВерсияПлатформы(v string, rac string, v8 string) *ВерсияПлатформы {
	return &ВерсияПлатформы{
		v,
		rac,
		v8,
	}
}

var доступныеВерсииПлатформы = make(map[string]*ВерсияПлатформы)

func init() {
	log.SetLevel(log.DebugLevel)
	log.Debugf("Текущая операционная система: %s", runtime.GOOS)
	собратьКешДоступныхВерсий()

}
func добавитьВерсию(v *ВерсияПлатформы) {
	доступныеВерсииПлатформы[v.Версия] = v
}

func ПолучитьСписокДоступныхВерсий() map[string]*ВерсияПлатформы {

	return доступныеВерсииПлатформы
}

func ПолучитьВерсию(строкаВерсияПлатформы string) (v *ВерсияПлатформы, err error) {

	if !strings.HasPrefix(строкаВерсияПлатформы, "8.") {
		log.Panicf("Неверная версия платформы < %s >", строкаВерсияПлатформы)
	}

	количествоТочекВЗапрошеннойВерсии := strings.Count(строкаВерсияПлатформы, ".")
	ИскомаяВерсия := строкаВерсияПлатформы
	if количествоТочекВЗапрошеннойВерсии < 3 {

		for версия := range доступныеВерсииПлатформы {
			if strings.HasPrefix(версия, строкаВерсияПлатформы) && version.Compare(ИскомаяВерсия, версия, "<") {

				log.Debugf("Найдена более старшая версия %s > %s", версия, ИскомаяВерсия)

				ИскомаяВерсия = версия

			}
		}
	}
	var ok bool
	v, ok = доступныеВерсииПлатформы[ИскомаяВерсия]

	if !ok {
		err = errors.Errorf("Запрошена не установленная версия платформы < %s >", строкаВерсияПлатформы)
	} else {
		log.Debugf("Использую версию %s", ИскомаяВерсия)
	}

	return
}

func ПолучитьВерсиюПоУмолчанию() (v *ВерсияПлатформы) {
	v, _ = ПолучитьВерсию(ВерсияПоУмолчанию)
	return
}

//noinspection ALL
func собратьКешДоступныхВерсий() {

	matchWindows := "\\*\\bin\\1cv8.exe"

	match1cv8 := "/*/1cv8"
	fRacName := "rac"

	var МассивПутей []string

	if этоWindows {

		МассивПутей = собратьВозможныеКаталогиУстановкиПлатформыWindows()
		match1cv8 = matchWindows
		fRacName = "rac.exe"

	} else {
		МассивПутей = собратьВозможныеКаталогиУстановкиПлатформыLinux()
	}

	if len(МассивПутей) == 0 {
		log.Debugf("Не обнаружено установленных версий платформы 1С")
		return
	}
	log.Debugf("Массив найденных путей установки 1С: %v", МассивПутей)

	for _, проверяемыйПуть := range МассивПутей {

		files, _ := filepath.Glob(проверяемыйПуть + match1cv8)

		log.Debugf("Нашли подходящих файлов: %v", cap(files))

		for _, fV8 := range files {

			var ВерсияПлатформы string

			fileDir := filepath.Dir(fV8)
			fRacName = filepath.Join(fileDir, fRacName)

			if !этоWindows {
				ВерсияПлатформы = получитьВерсиюПоRac(fRacName)

			} else {
				ВерсияПлатформы = получитьВерсиюИзПути(fV8)
			}

			if len(ВерсияПлатформы) == 0 {
				ВерсияПлатформы = ВерсияПоУмолчанию
			}

			добавитьВерсию(новаяВерсияПлатформы(ВерсияПлатформы, fRacName, fV8))

		}

	}

}

func получитьВерсиюИзПути(fV8 string) (version string) {

	regExpVersion := "\\d+(\\.\\d+)+"

	re := regexp.MustCompile(regExpVersion)

	version = re.FindString(fV8)

	return
}

func получитьВерсиюПоRac(pathRac string) (version string) {

	if _, err := v8tools.Exists(pathRac); err != nil {
		log.Debugf("Не удалось прочитать версию 1С по причине: %s ", err)
		return
	}

	out, execErr := exec.Command(pathRac, "-v").Output()
	if execErr != nil {
		log.Debugf("Не удалось прочитать версию 1С по причине %s", execErr)
		return
	}
	log.Debugf("Вывод команды rac -v: %s", out)

	version = strings.TrimSpace(string(out))

	return

}

//noinspection ALL
func собратьВозможныеКаталогиУстановкиПлатформыWindows() (МассивПутей []string) {

	var СуффиксРасположения = filepath.Join("1C", "1CEStart", "1CEStart.cfg")

	var envs = []string{
		"ALLUSERSPROFILE",
		"APPDATA",
	}

	for _, env := range envs {
		if cat, ok := os.LookupEnv(env); ok {
			дополнитьМассивРасположенийИзКонфигурационногоФайла(filepath.Join(cat, СуффиксРасположения), &МассивПутей)
		}
	}

	if len(МассивПутей) == 0 {
		log.Debugf("В конфигах стартера не найдены пути установки. Пробую стандартные пути наугад.")

		стандартныеПутиУстановки := []string{
			filepath.Join("C:", "Program Files (x86)"),
			filepath.Join("C:", "Program Files"),
		}

		for _, путьУстановки := range стандартныеПутиУстановки {
			if ok, _ := v8tools.Exists(путьУстановки); ok {
				МассивПутей = append(МассивПутей, путьУстановки)
			}
		}

	}

	return МассивПутей

}

func собратьВозможныеКаталогиУстановкиПлатформыLinux() (МассивПутей []string) {

	var КорневойПуть1С = filepath.Join("/opt", "1C", "v8.3")

	if ok, _ := v8tools.Exists(КорневойПуть1С); ok {

		МассивПутей = append(МассивПутей, КорневойПуть1С)
	}

	return
}

func дополнитьМассивРасположенийИзКонфигурационногоФайла(ИмяФайла string, МассивПутей *[]string) {

	log.Debugf("Читаю файл настроек: %s", ИмяФайла)

	Конфиг, err := КонфигурацияСтартера.ПрочитатьНастройкиСтартера(ИмяФайла)

	if err != nil {
		log.Errorf("Не удалось прочитать файл конфига стартера 1С: %s", err)
		return
	}

	var Значения = Конфиг.ПолучитьНастройку("InstalledLocation")

	log.Debugf("Начальное состояние МассивПутей: %s", МассивПутей)
	for _, item := range Значения {
		*МассивПутей = append(*МассивПутей, item)

		log.Debugf("Добавлен элемент %s к массиву", item)
	}
	log.Debugf("Конечное состояние МассивПутей: %s", МассивПутей)

}
