package logging

import (
	"fmt"
	"log"
	"os"
)

type LogPrinter interface {
	Printf(string, ...interface{})
	Print(...interface{})
}

type logger struct {
	on         bool
	loggerInfo LogPrinter
}

func (lg *logger) SetOn(on bool) *logger {
	lg.on = on
	return lg
}

func (lg *logger) Print(msg any, params ...any) {
	if !lg.on || msg == nil {
		return
	}
	switch msg := msg.(type) {
	case error:
		lg.loggerInfo.Print(msg.Error())
	case string:
		lg.loggerInfo.Printf(msg, params...)
	default:
		lg.loggerInfo.Printf(fmt.Sprintf("%v", msg), params...)
	}
}

func GetLogger(name string) *logger {
	return &logger{
		loggerInfo: log.New(os.Stdout, fmt.Sprintf("[ %s ] ", name), log.Ltime),
	}
}
