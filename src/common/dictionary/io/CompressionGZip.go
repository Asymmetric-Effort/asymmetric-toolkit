package io

/*
	Compression::Gzip() - Facilitates compression using the GZIP library.
 */
import (
	"bytes"
	gzipCompression "compress/gzip"
)

func (o *Compression) Gzip(in *[]byte) (out *[]byte) {
	var buf bytes.Buffer

	// Write gzipped data to the client
	g, err := gzipCompression.NewWriterLevel(&buf, gzipCompression.BestSpeed)
	g.Name = CompressionHeader
	if err != nil {
		panic(err)
	}

	_, err = g.Write(*in)
	if err != nil {
		panic(err)
	}

	err = g.Close()
	if err != nil {
		panic(err)
	}

	data := buf.Bytes()
	return &data
}
