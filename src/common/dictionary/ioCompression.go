package dictionary

type ioCompression uint8

const (
	noCompress = 0
	gzip       = 1
)

const (
	CompressionHeader = "Asymmetric-Toolkit(GZIP)"
)