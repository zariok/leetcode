package main

func minRemoveToMakeValid(s string) string {
	openCnt := 0
	closeCnt := 0
	newStr := ""

	for _, val := range s {
		if val == '(' {
			openCnt++
		} else if val == ')' {
			if openCnt < closeCnt+1 {
				continue
			}
			closeCnt++
		}
		newStr += string(val)
	}

	if openCnt > closeCnt {
		remove := openCnt - closeCnt
		for i := remove; i > 0; i-- {
			for i := len(newStr) - 1; i >= 0; i-- {
				if string(newStr[i]) == "(" {
					newStr = newStr[:i] + newStr[i+1:]
					break
				}
			}
		}
	}

	return newStr
}

/*
func main() {
	fmt.Printf("Returned: %s\n", minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Printf("Returned: %s\n", minRemoveToMakeValid("a)b(c)d"))
	fmt.Printf("Returned: %s\n", minRemoveToMakeValid("))(("))
	fmt.Printf("Returned: %s\n", minRemoveToMakeValid("())()((("))

}
*/
