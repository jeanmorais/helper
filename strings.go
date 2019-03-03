package helper

import (
	"regexp"
	"strings"
)

func OnlyLetters(str string) string {
	regx := regexp.MustCompilePOSIX("[a-zA-Z]+").FindAllString(str, -1)
	return strings.Join(regx[:], "")
}

func OnlyNumbers(numbers string) string {
	str := regexp.MustCompilePOSIX("[0-9]+").FindAllString(numbers, -1)
	return strings.Join(str[:], "")
}

func LeftPad(str string, length int, pad string) string {
	return formatPadding(pad, length-len(str)) + str
}

func RightPad(str string, length int, pad string) string {
	return str + formatPadding(pad, length-len(str))
}

func formatPadding(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}
