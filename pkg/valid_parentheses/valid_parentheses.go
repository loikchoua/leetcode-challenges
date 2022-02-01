package validparentheses

func IsValid(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			}
			if isCloseParentheses(stack[len(stack)-1], string(s[i])) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isCloseParentheses(open, end string) bool {
	switch {
	case open == "(" && end == ")":
		return true
	case open == "[" && end == "]":
		return true
	case open == "{" && end == "}":
		return true
	}
	return false
}
