package main

import (
	"blg/blg/db"
	"blg/blg/serve"
	"blg/tools/cnf"

	qlog "github.com/RRRRomeo/QLog/api"
)

func main() {
	cnf.ReadCnf()
	db.Init()
	if !serve.Init() {
		qlog.Errf("serve init fail!\n")
		return
	}

	if !serve.EventsHandler() {
		qlog.Errf("handle events fail\n")
		return
	}

	if !serve.Run(":54591") {
		qlog.Errf("serve run fail\n")
		return
	}
}
