package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func runeRepeat(forRepeat, count rune) string {
	repeat := string(forRepeat)
	c, _ := strconv.Atoi(string(count))
	return strings.Repeat(repeat, c)
}

func Unpack(s string) (string, error) {
	str := []rune(s)
	var (
		result strings.Builder
		next   bool
	)
	for i := 0; i < len(str); i++ {
		if next {
			next = false
			continue
		}
		fmt.Println(i, string(str[i]))
		if !unicode.IsLetter(str[i]) {
			return "", ErrInvalidString
		}
		if i+1 < len(str) && unicode.IsNumber(str[i+1]) {
			result.WriteString(runeRepeat(str[i], str[i+1]))
			next = true
			continue
		}
		result.WriteRune(str[i])
	}
	return result.String(), nil
}
