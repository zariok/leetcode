package main

func minRemoveToMakeValid2(s string) string {
	openCnt := 0
	closeCnt := 0
	newStr := make([]rune, 0)

	for _, val := range []rune(s) {
		if val == '(' {
			openCnt++
		} else if val == ')' {
			if openCnt < closeCnt+1 {
				continue
			}
			closeCnt++
		}
		newStr = append(newStr, val)
	}

	if openCnt > closeCnt {
		remove := openCnt - closeCnt
		for i := remove; i > 0; i-- {
			for i := len(newStr) - 1; i >= 0; i-- {
				if newStr[i] == '(' {
					newStr = append(newStr[:i], newStr[i+1:]...)
					break
				}
			}
		}

	}

	return string(newStr)
}
