package cli

import (
	"regexp"
)

func (o *Configuration) String(name string, defaultValue string, pattern string) ValidationFunc {
	re:=regexp.MustCompile(`.+`)
	return func(i *string) bool {
		return re.MatchString(*i)
	}
}
