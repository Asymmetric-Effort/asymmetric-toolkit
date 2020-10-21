package file_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
	"testing"
)

func TestDirExists(t *testing.T) {
	baseDir, _ :=os.Getwd()
	if !file.DirExists(baseDir){
		fmt.Println("baseDir:",baseDir)
		t.FailNow()
	}
}