package errors

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestFatal(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		Fatal(1, "expect true") //cause assertion error."
		return
	}
	prog:=os.Args[0]
	args:="-test.run=TestFatal"
	cmd := exec.Command(prog, args)
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	fmt.Println(prog, args)
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)
	if ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", e)
}
