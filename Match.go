package my

import (
	"strings"
)

//Match returns the index of the matching endparen/brace/bracket
func Match(str string, left rune) int {
	var right rune
	switch left {
	case '(':
		right = ')'
	case '{':
		right = '}'
	case '[':
		right = ']'
	case '<':
		right = '>'
	default:
		return -1
	}
	return MatchUneven(str, left, right)
}

func MatchUneven(str string, left, right rune) int {
	first := true
	count := 0
	return strings.IndexFunc(str, func(r rune) bool {
		if r == left {
			count++
			first = false
		}
		if r == right {
			count--
		}
		if count == 0 {
			if !first {
				return true
			}
			return false
		}
		return false
	})
}
