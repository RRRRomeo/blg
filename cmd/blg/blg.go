package main

import (
	"blg/blg/db"
	"blg/blg/serve"
	"blg/tools/cnf"
	"fmt"

	qlog "github.com/RRRRomeo/QLog/api"
)

func main() {
	cnf.ReadCnf()
	db.Init()
	sercnf := cnf.GlobalCnf.Server
	if !serve.Init() {
		qlog.Errf("serve init fail!\n")
		return
	}

	if !serve.EventsHandler() {
		qlog.Errf("handle events fail\n")
		return
	}

	ser := fmt.Sprintf("%s:%s", sercnf.Addr, sercnf.Port)
	fmt.Printf("ser:%s\n", ser)
	if !serve.Run(ser) {
		qlog.Errf("serve run fail\n")
		return
	}
}
