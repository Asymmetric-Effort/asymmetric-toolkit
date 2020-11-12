package logger


func (o *LogSettings) Set(key string, value string){
	(*o)[key] = value
}
