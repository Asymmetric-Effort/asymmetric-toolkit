package file_test

import (
	"fmt"
	"os"
	"testing"
)

func TestDirExists(t *testing.T) {
	baseDir, _ :=os.Getwd()
	if !DirExists(baseDir){
		fmt.Println("baseDir:",baseDir)
		t.FailNow()
	}
}