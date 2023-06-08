package main

import (
	"blg/blg/db"
	"blg/blg/serve"

	qlog "github.com/RRRRomeo/QLog/api"
)

func main() {
	qlog.Debugf("test\n")
	db.Init()
	if !serve.Init() {
		qlog.Errf("serve init fail!\n")
		return
	}

	if !serve.EventsHandler() {
		qlog.Errf("handle events fail\n")
		return
	}

	if !serve.Run2(":54591") {
		qlog.Errf("serve run fail\n")
		return
	}
}
