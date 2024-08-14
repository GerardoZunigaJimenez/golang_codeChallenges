package alphanumeric

import "regexp"

func RemoveNonAlphanumeric(s string) string {
	regex := regexp.MustCompile("[a-zA-Z0-9_]*")

	var r string
	for _, v := range regex.FindAllString(s, 10) {
		r += v
	}
	return r
}

func RemoveNonAlphanumeric2(s string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
}
