package kata

import (
	"strings"
)

func FindShort(s string) int {
	words := strings.Split(s, " ")
	ret := 1<<63 - 1

	for _, word := range words {
		if len(word) < ret {
			ret = len(word)
		}
	}

	return ret
}
