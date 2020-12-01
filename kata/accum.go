package kata

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	ret := ""
	s = strings.ToLower(s)
	for i, letter := range s {
		tmp := ""
		for j := 0; j < i+1; j++ {
			tmp += string(letter)
		}

		ret += fmt.Sprintf("%s-", strings.Title(tmp))
	}

	return ret[:len(ret)-1]
}
