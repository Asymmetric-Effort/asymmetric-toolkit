package dictionary

/*
	dictionary.File::Open() will open/create a dictionary file.  If the file doesn't exist, the file
						    will be created with a default header.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
)

func (o *File) Open(filename string) {
	errors.Assert(o.Handle == nil, "Cannot execute File::Open() when a file is already open.")
	//
	// Open the dictionary file.  If it does not exist, then create a
	// new dictionary file by calling File::create().
	//
	var newFile = !file.FileExists(filename)
	if newFile {
		o.Create(filename) //Create then close.
	}
	//
	// check our error state
	//
	errors.Assert(o.Err == nil, fmt.Sprintf("%v", o.Err))
	//
	// Open for read-write access
	//
	o.Handle, o.Err = os.OpenFile(filename, os.O_RDWR, filePermissions)
	errors.Assert(o.Err == nil, fmt.Sprintf("%v", o.Err))
	//
	// Load the header.
	//
	o.ReadHeader()
	errors.Assert(o.Header.FileVersion == fileVersion, "File version mismatch in dictionary file")
	errors.Assert(o.Header.ScoreVersion == scoreVersion, "Score version mismatch in dictionary file")
	errors.Assert(o.Err == nil, fmt.Sprintf("%v", o.Err))
	errors.Assert(o.Header != nil, "After loading the file header, the header should not be nil.")
	return
}
