package logger

import (
	"fmt"
	"strings"
)

type LogLevel int

const (
	Critical LogLevel = 0
	Error    LogLevel = 1
	Warning  LogLevel = 2
	Info     LogLevel = 3
	Debug    LogLevel = 4

	LevelStrings string = "critical,error,warning,info,debug"
)

func (o *LogLevel) Get() (v LogLevel) {
	switch *o {
	case Critical, Error, Warning, Info, Debug:
		v = *o
		break
	default:
		panic(fmt.Sprintf("Programming error: invalid Log level: %d", *o))
	}
	return
}

func (o *LogLevel) Set(v LogLevel) {
	switch v {
	case Critical, Error, Warning, Info, Debug:
		break
	default:
		panic(fmt.Sprintf("Programming error: invalid Log level: %d", v))
		*o = Debug
		return
	}
	*o = v
}

func (o *LogLevel) String() (s string) {
	list := strings.Split(LevelStrings, ",")
	i := int(o.Get())
	if i < len(list) {
		s=strings.ToLower(list[i])
	} else {
		panic(fmt.Sprintf("Programming error: Invalid Log level: %d", *o))
	}
	return
}
