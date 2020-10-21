package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func getTestFile() string {
	return filepath.Join(os.Args[0])
}

func TestFilePath_Exists_Happy(t *testing.T) {
	var fp types.FilePath = types.FilePath(getTestFile())
	errors.Assert(fp.Exists(),fmt.Sprintf("Expected file to exist: %s",fp))
}

func TestFilePath_Exists_BadFile(t *testing.T) {
	defer func(){recover()}()
	var fp types.FilePath = types.FilePath("non_existent_bad_file_that_should_never_exist.ext")
	errors.Assert(!fp.Exists(),fmt.Sprintf("Expected file not to exist: %s",fp))
}

func TestFilePath_SetGet(t *testing.T){
	var fp types.FilePath
	fp.Set(getTestFile())
	errors.Assert(fp.Get() == string(fp),"FilePath::Get() and ::Set() return different answers")
}

func TestFilePath_OpenRead(t *testing.T) {
	var fp types.FilePath = types.FilePath(getTestFile())
	file:=fp.OpenRead()
	defer func(){
		if err:=file.Close();err!=nil{
			panic(err)
		}
	}()
	errors.Assert(fp.Exists(),fmt.Sprintf("Expected file to exist: %s",fp))
}

func TestFilePath_OpenWrite(t *testing.T) {
	t.SkipNow()
	var fp types.FilePath = types.FilePath("./tempFile.txt")
	file:=fp.OpenWrite()
	defer func(){
		if err:=file.Close();err!=nil{
			panic(err)
		}
	}()
	errors.Assert(fp.Exists(),fmt.Sprintf("Expected file to exist: %s",fp))
}
