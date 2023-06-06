package log

import (
	qlog "github.com/RRRRomeo/QLog/api"
)

const (
	ROOTLOGFILEPATH string = ""
)

var BlgLogger = qlog.LoggerInit(0, 1, "/home/xs/data/code_space/blg/logs")
