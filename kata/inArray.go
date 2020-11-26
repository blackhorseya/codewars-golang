package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	sort.Strings(array1)

	ret := []string{}
	for _, a1 := range array1 {
		for _, a2 := range array2 {
			if strings.Contains(a2, a1) {
				if appendIf(ret, a1) {
					ret = append(ret, a1)
					break
				}
			}
		}
	}

	return ret
}

func appendIf(arr []string, target string) bool {
	for _, word := range arr {
		if word == target {
			return false
		}
	}

	return true
}
