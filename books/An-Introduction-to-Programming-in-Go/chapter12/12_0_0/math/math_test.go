package math

import "testing"

func TestAvg(t *testing.T) {
	var v float64
	v = Avg([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
