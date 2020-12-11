package dictionary
/*
	dictionary.File::Create() will create a new dictionary file with a default header struct.
 */
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
)

func (o *File) Create(filename string) {
	errors.Assert(o.Handle == nil, "Cannot execute File::create() when a file is already open.")
	//
	// Create a new file.
	//
	o.Handle, o.Err = os.Create(filename)
	errors.Assert(o.Err == nil, fmt.Sprintf("%v", o.Err))
	defer file.CloseFile(o.Handle)
	//
	// Write a blank header.
	//
	o.WriteHeader(&Header{
		fileVersion,
		scoreVersion,
		0, // We have no root node yet
		0, // We have no root tag yet.
	})
}
