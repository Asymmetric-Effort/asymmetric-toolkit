package persistentTree

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
	"testing"
)

func TestDeleteFile(t *testing.T) {

	t.Log("TestDeleteFile(): Start")

	const testFileName = "TestDeleteFile.txt"

	fh, err := os.Create(testFileName)
	if err != nil {
		panic(fmt.Sprintf("persistentTree::TestDeleteFile() create file error: %v", err))
	}
	t.Log("TestDeleteFile(): Created test file.")

	_, err = fh.Write([]byte("test"))
	if err != nil {
		panic(fmt.Sprintf("persistenTree::TestDeleteFile() file write error: %v", err))
	}
	t.Log("TestDeleteFile(): file write OK")

	err = fh.Close()
	if err != nil {
		panic(fmt.Sprintf("persistentTree::TestDeleteFile() close file error: %v", err))
	}
	t.Log("TestDeleteFile(): close file OK")

	deleteFile(testFileName)
	t.Log("TestDeleteFile(): OK (done)")

	defer func() {
		t.Log("Deferred cleanup function.")
		if !file.FileExists(testFileName) {
			return
		}
		err = os.Remove(testFileName)
		if err != nil {
			panic(err)
		}
	}()
	t.Log("TestDeleteFile(): Done")
}
