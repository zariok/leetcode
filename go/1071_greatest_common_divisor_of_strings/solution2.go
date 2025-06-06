package main

// Given two strings `str1` and `str2`, return the largest string `x` such that `x` divides both `str1` and `str2`.
func gcdOfStrings2(str1 string, str2 string) string {
	// get the smallest `x` to start
	x := str1
	if len(x) > len(str2) {
		x = str2
	}

	// iterate over `x`; check if x divides str1 && str2, if not pop a letter from the end of `x`
	for idx := len(x); idx > 0; idx-- {
		// if `x` is a divisor of both str1 and str2, return it
		if isDivisor(str1, x) && isDivisor(str2, x) {
			return x
		}
		// pop one letter from end of string
		x = string(x[:len(x)-1])
	}

	return ""
}

func isDivisor(str string, divisor string) bool {
	// if same, is divisor
	if str == divisor {
		return true
	}

	strLen := len(str)
	divLen := len(divisor)
	// sanity; does the length make it possible?
	if strLen%divLen != 0 {
		return false
	}

	// iterate while idx is < the string length; This looks at slices of the string of divLen
	for idx := 0; idx < strLen; idx += divLen {
		// get a slice of str from idx to index + length of divisor, if it does NOT
		// equal divisor, return false;
		if string(str[idx:idx+divLen]) != divisor {
			return false
		}
	}

	return true
}
