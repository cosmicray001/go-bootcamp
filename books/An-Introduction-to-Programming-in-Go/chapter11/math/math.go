package math

// Finds the average of a series of numbers
func Avg(nums []float64) float64 {
	var total float64 = 0.0
	for _, val := range nums {
		total += val
	}
	return total / float64(len(nums))
}