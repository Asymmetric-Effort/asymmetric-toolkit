package dictionary

import (
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
