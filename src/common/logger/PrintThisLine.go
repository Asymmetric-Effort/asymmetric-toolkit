package logger
/*
	Logger::PrintThisLine() is a function used to evaluate the line logging level against the Logger's current
	logging level to determine whether a given log line is actually printed to screen or just dropped.
 */
func (o *Logger) PrintThisLine(lineLevel Level) bool {
	return (lineLevel.Get() <= o.Level.Get())&& (o.Writer != nil)
}
