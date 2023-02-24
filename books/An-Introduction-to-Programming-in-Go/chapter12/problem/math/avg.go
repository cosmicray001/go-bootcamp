package math

func Avg(nums []float64) float64 {
	if len(nums) < 1 {
		return -1
	}
	total := float64(0)
	for _, val := range nums {
		total += val
	}
	return total / float64(len(nums))
}
