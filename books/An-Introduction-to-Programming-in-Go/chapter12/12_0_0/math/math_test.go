package math

import "testing"

type testpair struct {
	nums []float64
	avg  float64
}

var tests = []testpair{
	{nums: []float64{1, 2}, avg: 1.5},
	{nums: []float64{1, 1, 1, 1, 1, 1}, avg: 1},
	{nums: []float64{-1, 1}, avg: 0},
	{nums: []float64{2, 4}, avg: 3},
}

func TestAvg(t *testing.T) {
	for _, pair := range tests {
		val := Avg(pair.nums)
		if val != pair.avg {
			t.Error(
				"For", pair.nums,
				"expected", pair.avg,
				"got", val,
			)
		}
	}
}
