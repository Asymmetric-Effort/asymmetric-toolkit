package counter
/*
	Counter::Setup() initializes the Counter.  This test will run through the test and verify
	the operation.
 */
import (
	"testing"
)

func TestCounter_Setup(t *testing.T) {
	tests := []struct {
		charset  string
		wordSize int
		pass     bool
	}{
		{
			"",
			5,
			false,
		}, {
			"0123456789",
			-1,
			false,
		}, {
			"0123456789",
			0,
			false,
		}, {
			"0123456789",
			1,
			true,
		},{
			"0123456789",
			2,
			true,
		},
	}
	onErrorRecover := func() { recover() }
	onErrorFail := func() {}
	onError := onErrorRecover
	testFunc := func(charset string, wordSize int, pass bool) {
		var c Counter
		if pass {
			onError = onErrorFail
		} else {
			onError = onErrorRecover
		}
		defer onError()
		c.Setup(&charset, wordSize)
		if pass {
			//ToDo: verify operation in this test.
		}else{
			panic("Expected error when given empty charset.")
		}
	}

	for _, test:=range tests{
		testFunc(test.charset,test.wordSize,test.pass)
	}
}
