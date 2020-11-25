package kata

import (
	"strings"
)

func ToCamelCase(s string) string {
	toUpper := false
	var ret string
	for _, c := range s {
		if c == '-' || c == '_' {
			toUpper = true
			continue
		}

		if toUpper {
			upper := strings.ToUpper(string(c))
			ret += upper
			toUpper = false
		} else {
			ret += string(c)
		}
	}

	return ret
}
