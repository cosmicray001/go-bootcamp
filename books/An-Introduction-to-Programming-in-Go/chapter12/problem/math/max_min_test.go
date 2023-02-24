package math

import "testing"

type testpair struct {
	nums []float64
	v float64
}

var maxTests = []testpair{
	{[]float64{1, 2, 3, 4}, 4},
	{[]float64{100, 2, 3, 4}, 100},
	{[]float64{}, -1},
	{[]float64{1000}, 1000},
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		val := Max(pair.nums)
		if val != pair.v {
			t.Error(
				"for:", pair.nums,
				"got:", val,
				"expected:", pair.v,
			)
		}
	}
}

var minTests = []testpair{
	{[]float64{1, 2, 3, 4}, 1},
	{[]float64{100, 2, 3, 4, -1}, -1},
	{[]float64{}, -1},
	{[]float64{1000}, 1000},
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		val := Min(pair.nums)
		if val != pair.v {
			t.Error(
				"for:", pair.nums,
				"got:", val,
				"expected:", pair.v,
			)
		}
	}
}
