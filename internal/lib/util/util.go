package util

import "strings"

func ToResourceName(i string) string {
	r := strings.NewReplacer(".", "-", "*", "wildcard")
	return strings.ToLower(r.Replace(i))
}

func ToZoneName(i string) string {
	r := strings.NewReplacer(".", "_", "*", "wildcard")
	return strings.ToLower(r.Replace(i))
}