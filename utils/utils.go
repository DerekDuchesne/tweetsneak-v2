package utils

import "strings"

// StripChars will return a new string that is str without any instances of chars.
func StripChars(str, chars string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chars, r) < 0 {
			return r
		}
		return -1
	}, str)
}

// Min return the minimum value of x and y.
func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
