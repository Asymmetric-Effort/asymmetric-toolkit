package cli

type Expectation int

const (
	ExpectFlag  Expectation = 0
	ExpectValue Expectation = 1
)
