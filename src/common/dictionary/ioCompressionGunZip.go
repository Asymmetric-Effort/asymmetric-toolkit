package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	gzipCompression "compress/gzip"
	"fmt"
	"io/ioutil"
)

func (o *ioCompression) GunZip(in *[]byte) (out *[]byte) {
	var buf bytes.Buffer
	var data []byte

	g, err := gzipCompression.NewReader(bytes.NewBuffer(*in))
	if err != nil {
		panic(err)
	}
	errors.Assert(g.Name == CompressionHeader, CompressionHeader+" mismatch")

	data, err = ioutil.ReadAll(g)
	if err != nil {
		panic(fmt.Sprintf("Error(read):%v", err))
	}

	err = g.Close()
	if err != nil {
		panic(err)
	}

	buf.Write(data)
	data = buf.Bytes()
	return &data
}
