package math

func Max(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
	}
	return mx
}
