package cli

import (
	"strconv"
)

func (o *Configuration) Int(name string, defaultValue int, low int, high int) ValidationFunc {
	if low>high {
		panic("Internal Error: Expect low < high in valid.IntRange()")
	}
	return func(i *string) bool {
		if *i == "" {
			o.args[name] = defaultValue
		}
		n, err:=strconv.Atoi(*i)
		if err != nil {
			return false
		}
		if n >= low && n <= high {
			o.args[name] = n
		}
		return true
	}
}