package cli

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestArgument_List(t *testing.T) {
	func() { //simple test
		parser := ParserList(",")
		listString := "A,B,C,D"
		err, list := parser(&listString)
		if err != nil {
			panic(err)
		}
		assert.Error(len(list.Value) == 4, "Length mismatch")
	}()
}
