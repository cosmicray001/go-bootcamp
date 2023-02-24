package math

func Min(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}
	mn := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < mn {
			mn = nums[i]
		}
	}
	return mn
}
