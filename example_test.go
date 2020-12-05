package v8_test

// Output:
// Go is a general-purpose language designed with systems programming in mind.

import (
	"fmt"
	v8 "github.com/v8platform/v8"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ExampleNewTempDir() {
	content := []byte("temporary file's content")
	dir := v8.NewTempDir("", "example")

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}

func ExampleExecute() {

	where, err := v8.CreateTempInfobase()

	if err != nil {
		log.Fatal(err)
	}

	epfFilePath := "./user_any.epf"

	what := v8.Execute(epfFilePath)

	if err := v8.Run(where, what); err != nil {
		log.Fatal(err)
	}
}

func ExampleExecute_params() {

	where, err := v8.CreateTempInfobase()

	if err != nil {
		log.Fatal(err)
	}

	epfFilePath := "./path_to_epf.epf"

	what := v8.Execute(epfFilePath, map[string]string{
		"param1": "any_string",
		"param2": "any_string",
	})

	if err := v8.Run(where, what); err != nil {
		log.Fatal(err)
	}
}

func ExampleLoadCfg() {

	infobase := v8.NewFileIB("./infobase")
	//infobase := v8.NewServerIB("app", "demobase")

	what := v8.LoadCfg("./1cv8.cf")
	//what := v8.DumpCfg("./1cv8.cf)
	//what := v8.DumpIB("./1cv8.dt)

	err := v8.Run(infobase, what)
	//err := v8.Run(infobase, what)
	//err := v8.Run(infobase, what, v8.WithTimeout(1), v8.WithPath("path-to-exe"))
	//err := v8.Run(infobase, what, v8.WithCredentials("infobase-user","pwd"), v8.WithUnlockCode("123"))

	if err != nil {
		log.Fatal(err)
	}
}

func ExampleInfobase_file() {

	ib := &v8.Infobase{
		Connect:  v8.FilePath{File: "./infobase_path"},
		User:     "Admin",
		Password: "password",
	}

	connStr := ib.ConnectionString()

	fmt.Printf("ConnectionString: %s", connStr)

	// Output:
	// ConnectionString: /IBConnectionString File='./infobase_path';Usr=Admin;Pwd=password

}

func ExampleInfobase_ConnectionString_file() {

	ib := &v8.Infobase{
		Connect:  v8.FilePath{File: "./infobase_path"},
		User:     "Admin",
		Password: "password",
	}

	connStr := ib.ConnectionString()

	fmt.Printf("ConnectionString: %s", connStr)

	// Output:
	// ConnectionString: /IBConnectionString File='./infobase_path';Usr=Admin;Pwd=password

}

func ExampleInfobase_ConnectionString_server() {

	ib := &v8.Infobase{
		Connect:  v8.ServerPath{Server: "server", Ref: "ib_name"},
		User:     "Admin",
		Password: "password",
	}

	connStr := ib.ConnectionString()

	fmt.Printf("ConnectionString: %s", connStr)

	// Output:
	// ConnectionString: /IBConnectionString Srvr=server;Ref='ib_name';Usr=Admin;Pwd=password

}

func ExampleNewInfobase_from_path() {

	ib, err := v8.NewInfobase("./.github")

	if err != nil {
		log.Fatal(err)
	}

	connStr := ib.ConnectionString()

	fmt.Printf("ConnectionString: %s", connStr)

	// Output:
	// ConnectionString: /IBConnectionString File='./.github';

}

func ExampleNewInfobase_connect() {

	ib, err := v8.NewInfobase("File=./file_ib;Locale=ru_RU;Usr=User;Pwd=Password;Prmod=1;LicDstr=Y;")

	if err != nil {
		log.Fatal(err)
	}

	connStr := ib.ConnectionString()

	fmt.Printf("ConnectionString: %s", connStr)

	// Output:
	// ConnectionString: /IBConnectionString File='./file_ib';Usr=User;Pwd=Password;LicDstr=Y;Prmod=1;Locale=ru_RU

}

func ExampleDumpIB_file() {

	infobase, err := v8.NewInfobase("File=./file_ib;Locale=ru_RU;Usr=User;Pwd=Password;Prmod=1;LicDstr=Y;")

	if err != nil {
		log.Fatal(err)
	}
	//infobase := v8.NewServerIB("app", "demobase")

	what := v8.DumpIB("./1cv8.dt")
	//what := v8.DumpCfg("./1cv8.cf)
	//what := v8.DumpIB("./1cv8.dt)

	err = v8.Run(infobase, what)
	//err := v8.Run(infobase, what)
	//err := v8.Run(infobase, what, v8.WithTimeout(1), v8.WithPath("path-to-exe"))
	//err := v8.Run(infobase, what, v8.WithCredentials("infobase-user","pwd"), v8.WithUnlockCode("123"))

	if err != nil {
		log.Fatal(err)
	}
}

func ExampleDumpIB_server() {

	infobase := v8.NewServerIB("app", "demobase")

	what := v8.DumpIB("./1cv8.dt")
	//what := v8.DumpCfg("./1cv8.cf)
	//what := v8.DumpIB("./1cv8.dt)

	err := v8.Run(infobase, what)
	//err := v8.Run(infobase, what)
	//err := v8.Run(infobase, what, v8.WithTimeout(1), v8.WithPath("path-to-exe"))
	//err := v8.Run(infobase, what, v8.WithCredentials("infobase-user","pwd"), v8.WithUnlockCode("123"))

	if err != nil {
		log.Fatal(err)
	}
}

func ExampleRestoreIB_file() {

	infobase, err := v8.CreateTempInfobase()
	if err != nil {
		log.Fatal(err)
	}

	//infobase := v8.NewServerIB("app", "demobase")

	what := v8.RestoreIB("./1cv8.dt")
	//what := v8.DumpCfg("./1cv8.cf)
	//what := v8.DumpIB("./1cv8.dt)

	err = v8.Run(infobase, what)
	//err := v8.Run(infobase, what)
	//err := v8.Run(infobase, what, v8.WithTimeout(1), v8.WithPath("path-to-exe"))
	//err := v8.Run(infobase, what, v8.WithCredentials("infobase-user","pwd"), v8.WithUnlockCode("123"))

	if err != nil {
		log.Fatal(err)
	}
}

func ExampleRun() {

	infobase, err := v8.CreateTempInfobase()
	if err != nil {
		log.Fatal(err)
	}

	//infobase := v8.NewServerIB("app", "demobase")

	what := v8.RestoreIB("./1cv8.dt")
	//what := v8.DumpCfg("./1cv8.cf)
	//what := v8.DumpIB("./1cv8.dt)

	err = v8.Run(infobase, what)
	//err := v8.Run(infobase, what)
	//err := v8.Run(infobase, what, v8.WithTimeout(1), v8.WithPath("path-to-exe"))
	//err := v8.Run(infobase, what, v8.WithCredentials("infobase-user","pwd"), v8.WithUnlockCode("123"))

	if err != nil {
		log.Fatal(err)
	}
}

func ExampleRun_with_opts() {

	infobase, err := v8.CreateTempInfobase()
	if err != nil {
		log.Fatal(err)
	}

	//infobase := v8.NewServerIB("app", "demobase")

	what := v8.RestoreIB("./1cv8.dt")
	//what := v8.DumpCfg("./1cv8.cf)
	//what := v8.DumpIB("./1cv8.dt)

	err = v8.Run(infobase, what,
		v8.WithTimeout(1), // указание таймаута выполнения в секундах
	//v8.WithPath("path-to-exe"), // указание пути к исполняемому файлу 1С.Предприятие
	//v8.WithUnlockCode("123"), // указание кода блокировки информационной базы
	//v8.WithDumpResult("./dump_result.txt"), // указание файла результата выполенния операции
	//v8.WithOut("./out.txt", false), // указание файла в который будет записан вывод консоли 1С.Предприятие
	//v8.WithVersion("8.3.16"), // Указание конкретной версии. Не работает с опцией v8.WithPath
	//v8.WithCredentials("Администратор", ""), // Указание пользователя и пароля для информационной базы
	//v8.WithCommonValues("/Visible", "/DisableStartupDialogs"), // Указание дополнительных опций запуска
	)

	if err != nil {
		log.Fatal(err)
	}
}
