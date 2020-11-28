package dictionary

/*
	Descriptor - This struct is the runtime representation of a dictionary file.  It also acts as the nexus
 				 for the methods used to manage the dictionary.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/definition"
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/headers"
)

type Descriptor struct {
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
