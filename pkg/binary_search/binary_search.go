package binarysearch

func Search(nums []int, target int) int {
	half := len(nums) / 2
	halfElement := nums[half]
	if halfElement == target {
		return half
	} else {
		if half == 0 {
			return -1
		} else if target < halfElement {
			return Search(nums[:half], target)
		} else {
			index := Search(nums[half:], target)
			if index != -1 {
				return half + index
			}
			return index
		}
	}
}

func SearchMoreEfficient(nums []int, target int) int {
	middle := 0
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle = left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
