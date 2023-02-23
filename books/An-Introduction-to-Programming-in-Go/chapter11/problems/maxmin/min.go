package maxmin

// This function return the min value from a list
func Min(nums []float64) float64 {
	if len(nums) == 0 {
		return -1.0
	}
	mn := nums[0]

	for i := 1; i < len(nums); i++ {
		if mn > nums[i] {
			mn = nums[i]
		}
	}
	return mn
}