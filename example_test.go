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

	where, err := v8.NewTempIB()
	if err != nil {
		log.Fatal(err)
	}
	epfFilePath := "./user_any.epf"

	what := v8.Execute(epfFilePath)

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
