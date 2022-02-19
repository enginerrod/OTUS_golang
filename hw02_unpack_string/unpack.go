package hw02unpackstring

import (
	"errors"
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
	var result string
	var next, escape = false, false
	if len(str) > 0 && !(unicode.IsLetter(str[0]) || str[0] == []rune(`\`)[0]) {
		return "", ErrInvalidString
	}
	for i, r := range str {
		switch {
		case next:
			next = false
			continue
		case !(unicode.IsNumber(r) || unicode.IsLetter(r) || r == []rune(`\`)[0]):
			return "", ErrInvalidString
		case r == []rune(`\`)[0] && !escape && i+1 < len(s):
			escape = true
			continue
		case escape && !(r == []rune(`\`)[0] || unicode.IsNumber(r)):
			return "", ErrInvalidString
		case i+1 < len(s):
			if unicode.IsLetter(r) && !unicode.IsNumber(str[i+1]) {
				result += string(r)
				continue
			}
			if unicode.IsLetter(r) && unicode.IsNumber(str[i+1]) {
				result += runeRepeat(r, str[i+1])
				next = true
				continue
			}
			if escape && (unicode.IsNumber(r) || r == []rune(`\`)[0]) && !unicode.IsNumber(str[i+1]) {
				result += string(r)
				escape = false
				continue
			}
			if escape && unicode.IsNumber(r) && unicode.IsNumber(str[i+1]) {
				result += runeRepeat(r, str[i+1])
				escape = false
				next = true
				continue
			}

			if escape && r == []rune(`\`)[0] && unicode.IsNumber(str[i+1]) {
				result += runeRepeat(r, str[i+1])
				escape = false
				next = true
				continue
			}

			if unicode.IsNumber(r) {
				return "", ErrInvalidString

			}

		default:
			if escape && r == []rune(`\`)[0] {
				result += string(r)
			}
			if !escape && (unicode.IsNumber(r) || r == []rune(`\`)[0]) {
				return "", ErrInvalidString
			}
			result += string(r)
		}
	}
	return result, nil
}
