package logger

import (
	"fmt"
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestLogger_TagClose(t *testing.T) {
	/*
		Setup the tests...
	*/
	var o TagTracker

	tests := []struct {
		tagName string
		tagId   TagId
	}{
		{
			"test1", // OK
			0,
		}, {
			"test2", // OK
			0,
		}, {
			"test3", // OK
			0,
		}, {
			"test4", // OK
			0,
		}, {
			"test5", // OK
			0,
		}, {
			"test6", // OK
			0,
		}, {
			"test7", // OK
			0,
		},
	}

	for i, test := range tests {
		test.tagId = o.Create(test.tagName)
		fmt.Printf("\ncreate: i:%2d tagId:%04d tagName:%-8.8s", i, test.tagId, test.tagName)
		o.Create(test.tagName)
	}
	/*
		Test the delete (close) operations.
	*/
	func() {
		startingCount:=len(tests)
		func() {
			o.Close(0)
			fmt.Println("Close tagId: 0 (ok)")
			defer func() { recover() }()
			o.Close(0)
			fmt.Println("expected error when closing the same tagId (0) twice.")
		}()
		fmt.Println("Close() test worked.")
		endingCount:=len(tests)
		assert.Error((endingCount - startingCount) == 1, "expected only one record to be deleted.")
	}()

}
