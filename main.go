package main

import (
	"fmt"

	validparentheses "github.com/loikchoua/leetcode-challenges/pkg/valid_parentheses"
)

func main() {
	s := "()"
	isValid := validparentheses.IsValid(s)
	fmt.Println(isValid)
}
