package dictionary

import (
	"encoding/binary"
	"os"
)

type Dictionary struct {
	//
	// Private properties
	//
	Config     *Configuration
	fileHandle *os.File
	counter    uint32 //Position counter (definition index).
	//
	//
	//
	Content struct {
		Header  HeaderDescriptor
		Content []DefinitionDescriptor
	}
}


func (o *Dictionary) WriteHeader(formatVersion Version, scoreVersion Version) {
	err := binary.Write(o.fileHandle, binary.LittleEndian, &o.Content.Header)
	if err != nil {
		panic(err)
	}
}

func (o *Dictionary) WriteDefinitions() {

}
func (o *Dictionary) LoadHeader() (err error) {
	err = binary.Read(o.fileHandle, binary.LittleEndian, &o.Content.Header)
	return
}

func (o *Dictionary) WriteDefinition(def Definition) {

}
func (o *Dictionary) ReadDefinition() *Definition {
	// Return nil if at EOF.
	return nil
}
