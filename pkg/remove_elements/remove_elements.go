package removeelements

func RemoveElements(nums []int, val int) int {
	count := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[j] == val {
			count++
			nums = append(nums[:j], nums[j+1:]...)
			nums = append(nums, 0)
		} else {
			j++
		}
	}
	return len(nums) - count
}
