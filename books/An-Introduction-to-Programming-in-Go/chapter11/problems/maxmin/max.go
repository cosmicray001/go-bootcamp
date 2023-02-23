package maxmin


// This function return the max value from a list 
func Max(nums []float64) float64 {
	if len(nums) == 0 {
		return -1.0
	}
	mx := nums[0]

	for i := 1; i < len(nums); i++ {
		if mx < nums[i] {
			mx = nums[i]
		}
	}
	return mx
}