package main

func sum(nums ...int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}

	return result
}
