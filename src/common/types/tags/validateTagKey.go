package tags

import "regexp"

func validateTag(key string) bool {
	re := regexp.MustCompile("[a-zA-Z][a-zA-Z0-9._]*[a-zA-Z0-9]+")
	if re.MatchString(key) {
		return true
	}
	panic("Invalid key in string tag.")
}
