package base7

import "fmt"

func ConvertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}
	resp := ""
	for num > 0 {
		resp = fmt.Sprintf("%d%s", num%7, resp)
		num /= 7
	}
	if negative {
		resp = "-" + resp
	}
	return resp
}

func getNumberOfDigits(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}
