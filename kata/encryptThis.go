package kata

import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	if len(text) == 0 {
		return ""
	}

	words := strings.Split(text, " ")

	var encode []string
	for _, word := range words {
		tmp, new := "", ""
		for i, char := range word {
			if i == 0 {
				new += strconv.Itoa(int(char))
			} else if i == 1 {
				tmp = string(char)
				new += string(word[len(word) - 1])
			} else if i == len(word) - 1 {
				new += tmp
			} else {
				new += string(char)
			}
		}
		encode = append(encode, new)
	}

	return strings.Join(encode, " ")
}
