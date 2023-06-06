package main

import (
	"blg/blg/db"

	qlog "github.com/RRRRomeo/QLog/api"
)

func main() {
	qlog.Debugf("test\n")
	db.Init()
}
