package file_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExists_File(t *testing.T) {
	baseDir, _ :=os.Getwd()
	f:=filepath.Join(baseDir,"exists_test.go")
	if !file.Exists(f,true){
		fmt.Println("file:",baseDir)
		t.FailNow()
	}
}

func TestExists_Directory(t *testing.T) {
	baseDir, _ :=os.Getwd()
	if !file.Exists(baseDir,false){
		fmt.Println("baseDir:",baseDir)
		t.FailNow()
	}
}