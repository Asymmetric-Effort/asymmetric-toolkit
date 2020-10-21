package file_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	baseDir, _ :=os.Getwd()
	if !file.ExistsAsFile(filepath.Join(baseDir,"ExistsAsFile_test.go")){
		t.FailNow()
	}
}