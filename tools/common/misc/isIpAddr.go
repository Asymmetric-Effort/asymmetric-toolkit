package misc

import (
	"regexp"
)

func IsIpAddr(s string) bool {
	const ipv4Pattern = `^(([01]{0,1}[0-9]{1,2}|2[0-4][0-9]|25[0-5])\.){3}(([01]{0,1}[0-9]{1,2}|2[0-4][0-9]|25[0-5])){1}$`
	re := regexp.MustCompile(ipv4Pattern)
	return re.MatchString(s)
}
