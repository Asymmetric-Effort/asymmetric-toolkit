package logger

/*
func TestLogWriterStdOut(t *testing.T) {
	var log Logger
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Printf(Debug, "%s", "Test")
	})
	msg := strings.Split(out, ":")
	word := strings.TrimSpace(msg[len(msg)-1])
	errors.Assert(word == "Test", fmt.Sprintf("Expected 'Test'. Encountered: '%s'", word))
}
*/