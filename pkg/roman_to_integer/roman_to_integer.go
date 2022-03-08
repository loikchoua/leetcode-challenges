package romantointeger

func RomanToInt(s string) int {
	convertMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := convertMap[string(s[0])]
	for i := 1; i < len(s); i++ {
		cur := convertMap[string(s[i])]
		prev := convertMap[string(s[i-1])]
		if cur <= prev {
			result += cur
		} else {
			result += cur - prev*2
		}
	}
	return result
}
