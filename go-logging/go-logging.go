//https://godoc.org/github.com/op/go-logging
//https://github.com/op/go-logging
package main

import (
	"github.com/op/go-logging"
	//"os"
)

var log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}")

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {
	//backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	//backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	//backend2Formatter := logging.NewBackendFormatter(backend2, format)
	//backend1Formatter := logging.NewBackendFormatter(backend1, format)

	//backend1Leveled := logging.AddModuleLevel(backend1)
	//backend1Leveled.SetLevel(logging.ERROR, "")

	//logging.SetBackend(backend1Leveled, backend2Formatter)
	//logging.SetBackend(backend1Leveled, backend1Formatter)
	//logging.SetBackend(backend1Leveled, nil)

	logging.SetLevel(logging.ERROR, "")
	logging.SetFormatter(format)

	log.Debugf("debug %s", Password("secret"))
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}
