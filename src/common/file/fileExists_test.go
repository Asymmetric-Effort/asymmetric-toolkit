package file_test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	baseDir, _ :=os.Getwd()
	if !FileExists(filepath.Join(baseDir,"fileExists_test.go")){
		t.FailNow()
	}
}