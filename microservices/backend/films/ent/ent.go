package ent

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[^a-z0-9]+")

func newSlug(s string) string {
	return strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
