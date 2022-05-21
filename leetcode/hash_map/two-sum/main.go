package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if diff == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
