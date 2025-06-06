package main

import (
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {

	tmpStr := ""
	gcdStr := ""
	for _, char := range str1 {
		tmpStr += string(char)
		if !strings.Contains(str2, tmpStr) {
			break
		}
		if len(strings.Replace(str1, tmpStr, "", -1)) == 0 && len(strings.Replace(str2, tmpStr, "", -1)) == 0 {
			gcdStr = tmpStr
		}
	}

	return gcdStr
}
