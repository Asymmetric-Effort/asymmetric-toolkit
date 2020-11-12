package logger

import "strconv"

func (o *LogSettings) GetBool(key string) bool{
	v,err:=strconv.ParseBool((*o)[key])
	if err!= nil {
		panic("LogSettings::SetInt expected boolean.  Encountered non-boolean input")
	}
	return v
}