package my

import (
	"strings"
)

func SplitAll(s string, sep string) []string {
	str := strings.Split(s, string(sep[0]))
	for i, se := range sep {
		se := string(se)
		if i == 0 {
			continue
		}
		for _, str2 := range str {
			str = append(str, strings.Split(str2, se)...)
		}
	}
	return str
}
