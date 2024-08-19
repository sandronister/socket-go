package utils

import "strings"

func PadLeft(str, pad string, length int) string {
	if len(str) >= length {
		return str
	}
	padding := strings.Repeat(pad, length-len(str))
	return padding + str
}
