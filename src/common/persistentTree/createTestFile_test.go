package persistentTree

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"testing"
)

func TestCreateTestFile(t *testing.T) {
	const testFileName = "TestCreateTestFile.txt"
	createTestFile(testFileName)

	defer deleteFile(testFileName)

	if file.FileExists(testFileName) {
		return
	}
	panic("TestCreateFile(): createTestFile() failed.")
}
