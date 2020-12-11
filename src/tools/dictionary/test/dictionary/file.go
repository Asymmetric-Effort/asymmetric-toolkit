package dictionary
/*

 */
import (
	"os"
	"sync"
)

type File struct {
	Handle   *os.File
	Err      error
	Header   *Header
	mutex    sync.Mutex
}
