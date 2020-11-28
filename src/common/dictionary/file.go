package dictionary

import "os"

type File struct {
	config *Configuration
	handle os.File
	position uint32
}

func (o *File) Setup(config *Configuration){
	o.config = config
	o.handle = os.OpenFile(config.FileName)
}

func (o *File) Teardown(config *Configuration){
	err:=o.handle.Close()
	if err != nil {
		panic(err)
	}
}