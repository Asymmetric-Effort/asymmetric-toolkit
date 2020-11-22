package io
/*
	Dictionary compression facility.
 */
type Compression uint8

const (
	NoCompress = 0
	Gzip       = 1
)

const (
	CompressionHeader = "Asymmetric-Toolkit(GZIP)"
)
