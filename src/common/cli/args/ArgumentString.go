package cliargs

/*
	The ArgumentString struct implements the Arguments interface{}
	and decodes the string cli parameter into the string value it
	represents.
*/
type ArgumentString struct {
	data string
}

func (o *ArgumentString) Get() string {
	return o.data
}

func (o *ArgumentString) Set(n string) {
	o.data = n
}

func (o *ArgumentString) SetString(s string, defaultValue string) {
	if s == "" {
		o.Set(defaultValue)
	} else {
		o.Set(s)
	}
}
