package palidromenumber

import "strconv"

func IsPalidrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	forward := []string{}
	backward := []string{}
	s := strconv.Itoa(x)
	for i := 0; i < len(s); i++ {
		forward = append(forward, string(s[i]))
		backward = append(backward, string(s[len(s)-1-i]))
	}
	for i := 0; i < len(forward); i++ {
		if forward[i] != backward[i] {
			return false
		}
	}
	return true
}

func IsPalidromeOptimized(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}
