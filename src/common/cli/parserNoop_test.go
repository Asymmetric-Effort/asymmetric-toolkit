package cli

import "testing"

func TestParserNoop(t *testing.T){
	p:=ParserNoop()
	err,val:=p(nil)
	if err != nil {
		panic(err)
	}
	if val != nil {
		panic("Noop should return nil value.")
	}
}