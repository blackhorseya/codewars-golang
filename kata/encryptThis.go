package kata

import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	ret := ""
	for _, word := range strings.Split(text, " ") {
		for i, letter := range word {
			switch i {
			case 0:
				ret += strconv.Itoa(int(letter))
			case 1:
				ret += string(word[len(word) - 1])
			case len(word) - 1:
				ret += string(word[1])
			default:
				ret += string(letter)
			}

		}
		ret += " "
	}

	return ret[:len(ret)-1]
}
