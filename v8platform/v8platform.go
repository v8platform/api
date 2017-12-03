package v8platform

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"../v8config-file"
	"../v8tools"
	log "github.com/sirupsen/logrus"
	version "github.com/mcuadros/go-version"
)

const ВерсияПоУмолчанию string = "8.3"
const этоWindows bool = runtime.GOOS == "windows"

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
	собратьКешДоступныхВерсий()
}
func добавитьВерсию(v *ВерсияПлатформы) {
	доступныеВерсииПлатформы[v.Версия] = v
}

func ПолучитьВерсию(строкаВерсияПлатформы string) (v *ВерсияПлатформы) {
	if !strings.HasPrefix(строкаВерсияПлатформы, "8.") {
		log.Panicf("Неверная версия платформы < %s >", строкаВерсияПлатформы)
	}

	version.Normalize(строкаВерсияПлатформы)

	v = доступныеВерсииПлатформы[строкаВерсияПлатформы]
	return
}

func ПолучитьВерсиюПоУмолчанию() (v *ВерсияПлатформы) {
	v = доступныеВерсииПлатформы[ВерсияПоУмолчанию]
	return
}

//noinspection ALL
func собратьКешДоступныхВерсий() {

	matchWindows := "\\1*8*\\*\\bin\\1cv8.exe"

	match1cv8 := "/*/1cv8"
	fRac := "rac"

	var МассивПутей []string

	if этоWindows {

		МассивПутей := собратьВозможныеКаталогиУстановкиПлатформыWindows()
		match1cv8 = matchWindows
		fRac = "rac.exe"

	} else {
		МассивПутей := собратьВозможныеКаталогиУстановкиПлатформыLinux()
	}

	if len(МассивПутей) == 0 {
		log.Debugf("Не обнаружено установленных версий платформы 1С")
		return
	}
	log.Debugf("Массив найденных путей установки 1С: %v", МассивПутей)

	for _, проверяемыйПуть := range МассивПутей {

		files, _ := filepath.Glob(srcDir + match1cv8)

		for _, fV8 := range files {

			var ВерсияПлатформы string

			fileDir = path.Dir(fV8)
			fRac = path.Join(fileDir, "rac.exe")

			if !этоWindows {
				ВерсияПлатформы = получитьВерсиюПоRac(fRac)

			} else {
				ВерсияПлатформы = получитьВерсиюИзПути(fV8)
			}

			if len(ВерсияПлатформы) == 0 {
				ВерсияПлатформы = ВерсияПоУмолчанию
			}

			добавитьВерсию(новаяВерсияПлатформы(Версия, fRac, fV8))

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

	var СуффиксРасположения = path.Join("1C", "1CEStart", "1CEStart.cfg")

	var envs = []string{
		"ALLUSERSPROFILE",
		"APPDATA",
	}

	for _, env := range envs {
		if cat, ok := os.LookupEnv(env); ok {
			дополнитьМассивРасположенийИзКонфигурационногоФайла(path.Join(cat, СуффиксРасположения), &МассивПутей)
		}
	}

	if len(МассивПутей) == 0 {
		log.Debugf("В конфигах стартера не найдены пути установки. Пробую стандартные пути наугад.")

		стандартныеПутиУстановки := []string{
			path.Join("C:", "Program Files (x86)"),
			path.Join("C:", "Program Files"),
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

	var КорневойПуть1С = path.Join("/opt", "1C", "v8.3")

	if ok, _ := v8tools.Exists(КорневойПуть1С); ok {

		МассивПутей = append(МассивПутей, КорневойПуть1С)
	}

	return
}

func дополнитьМассивРасположенийИзКонфигурационногоФайла(ИмяФайла string, МассивПутей *[]string) {

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
