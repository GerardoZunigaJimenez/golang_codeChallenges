package stringparser

import (
	"regexp"
	"strings"
)

func FindStringsUnderSlice(s string) []string {
	p := regexp.MustCompile(`\[(.*?)\]`).FindString(s)
	ip := strings.Split(strings.ReplaceAll(strings.Trim(p, "[]"), "\"", ""), ",")
	return ip
}
