package file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExists_File(t *testing.T) {
	baseDir, _ :=os.Getwd()
	f:=filepath.Join(baseDir,"exists_test.go")
	if !Exists(f,true){
		fmt.Println("file:",baseDir)
		t.FailNow()
	}
}

func TestExists_Directory(t *testing.T) {
	baseDir, _ :=os.Getwd()
	if !Exists(baseDir,false){
		fmt.Println("baseDir:",baseDir)
		t.FailNow()
	}
}