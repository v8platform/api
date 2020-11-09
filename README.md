go-v8platform
=============

[![ISC License](http://img.shields.io/badge/license-BSD-blue.svg)](http://copyfree.org)  

Реализация программного интерфейсом пакетного режима работы с 1С:Предприятием 8.x

## Пример работы
```go
package main

import "github.com/v8platform/v8"

func main() {
	
  infobase := v8.NewFileIB("./infobase")
  //infobase := v8.NewServerIB("app", "demobase")
  
  what := v8.LoadCfg("./1cv8.cf)
  //what := v8.DumpCfg("./1cv8.cf)
  //what := v8.DumpIB("./1cv8.dt)
  
  err := v8.Run(infobase, what)
  //err := v8.Run(infobase, what)
  //err := v8.Run(infobase, what, v8.WithTimeout(1), v8.WithPath("path-to-exe"))
  //err := v8.Run(infobase, what, v8.WithCredentials("infobase-user","pwd"), v8.WithUnlockCode("123"))
  
  if err != nil {
     println(err.Error())
  }
  
}
```
## Документация

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/v8platform/v8)
