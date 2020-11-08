package main_test

/*
	This is the top-level test program for testing CliDev.
*/

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
