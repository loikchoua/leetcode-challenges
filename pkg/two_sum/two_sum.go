package twosum

func TwoSum(nums []int, target int) []int {
	positionAndValues := map[int]int{}
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		positionAndValues[val] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		val, ok := positionAndValues[complement]
		if ok && val != i {
			return []int{i, val}
		}
	}
	return nil

}
