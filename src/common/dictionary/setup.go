package dictionary

/*
	Dictionary::Setup() will initialize the configuration object by consuming the object passed by caller
	and mapping it by reference to the internal pointer.  Then we open the dictionary file specified in the
	Configuration, load its header and setup our encryption/compression layer.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
)

func (o *Dictionary) Setup(config *Configuration) (err error) {
	o.Config = config //Load configuration from parent.

	if o.fileHandle != nil {
		return fmt.Errorf("File handle is already open.  Cannot open for read.")
	}
	if file.FileExists(o.Config.FileName) {
		if o.Config.Overwrite {
			err := os.Remove(o.Config.FileName)
			if err != nil {
				panic(err)
			}
		} else {
			return fmt.Errorf("file exists but no overwrite command was passed")
		}
	} // else file does not exist, we may create it.
	o.fileHandle, err = os.OpenFile(o.Config.FileName, FileAccess, FilePermissions)
	if err != nil {
		return err
	}
	err = o.LoadHeader() // This should leave our file handle at the first definition.
	return
}
