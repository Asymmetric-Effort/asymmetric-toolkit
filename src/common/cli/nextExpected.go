package cli

type NextExpected int

const (
	ExpectFlag  NextExpected = 0
	ExpectValue NextExpected = 1
	ExpectEnd   NextExpected = 2
	ExpectNone  NextExpected = 3
)