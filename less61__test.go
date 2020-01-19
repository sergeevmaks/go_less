package statistic

import "testing"

type testpair struct {
	values []float64
	summ   float64
}

var tests = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestSumm(t *testing.T) {
	for _, pair := range tests {
		v := Summ(pair.values)
		if v != pair.summ {
			t.Error(
				"For", pair.values,
				"expected", pair.summ,
				"got", v,
			)
		}
	}
}
