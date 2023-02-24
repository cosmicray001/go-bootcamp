package math

import "testing"

func TestAvg(t *testing.T) {
	nums := []float64{-1, 1}
	avg := Avg(nums)
	if avg != 0 {
		t.Error(
			"For:", nums,
			"got:", avg,
			"expected:", 0,
		)
	}
}