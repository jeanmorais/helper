package helper

import (
	"regexp"
	"strings"
)

// OnlyLetters returns the letters contained in str.
func OnlyLetters(str string) string {
	regx := regexp.MustCompilePOSIX("[a-zA-Z]+").FindAllString(str, -1)
	return strings.Join(regx[:], "")
}

// OnlyNumbers returns the numbers contained in str.
func OnlyNumbers(str string) string {
	regx := regexp.MustCompilePOSIX("[0-9]+").FindAllString(str, -1)
	return strings.Join(regx[:], "")
}

// LeftPad adds a "padding" to the left side of the string with the value pad.
func LeftPad(str string, length int, pad string) string {
	return formatPadding(pad, length-len(str)) + str
}
// RightPad adds a "padding" to the right side of the string with the value pad.
func RightPad(str string, length int, pad string) string {
	return str + formatPadding(pad, length-len(str))
}

func formatPadding(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}

// Reverse returns a string with the bytes of s in reverse order.
func Reverse(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}