package dictionary

/*
	Descriptor - This struct is the runtime representation of a dictionary file.  It also acts as the nexus
 				 for the methods used to manage the dictionary.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/headers"
	"os"
)

type Descriptor struct {
	/*
		dictionary.Configuration is the state passed in by the commandline argument.
			This runtime configuration is not persisted to disk.
	*/
	config *Configuration
	/*
		handle is the file handle for the dictionary file on disk.  Because we expect dictionary files
		can be very large, we keep the file handle open and read/write the file as needed.
	*/
	handle *os.File
	/*
		position is the current dictionary definition position (record number) in the dictionary file.
		This indicates current position and will be updated as the dictionary read/write operations move
		forward or backward.
	*/
	position uint32
	/*
		The HeaderDescriptor is the on-disk dictionary header.
	*/
	Header headers.Descriptor
	/*
		The DefinitionDescriptor is the on-disk cleartext current dictionary record.  As we navigate through
		the dictionary, the position property updates to indicate which record this descriptor represents.
	*/
	Definition definition.Descriptor
}
