package math

import "testing"

func TestAverage_Fails(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 2.5 {
		t.Error("Expected 2.5, got", v)
	}
}

// func TestAverage_Passes(t *testing.T) {
// 	var v float64
// 	v = Average([]float64{1, 2, 3, 4})
// 	if v != 2.5 {
// 		t.Error("Expected 2.5, got", v)
// 	}
// }

// type testpair struct {
// 	values  []float64
// 	average float64
// }

//// Execute many tests.
// var tests = []testpair{
// 	{[]float64{1, 2}, 1.5},
// 	{[]float64{1, 1, 1, 1, 1, 1}, 1},
// 	{[]float64{-1, 1}, 0},
// 	//{[]float64{1, 2, 3, 4}, 1.5},
// }

// func TestAverage(t *testing.T) {
// 	for _, pair := range tests {
// 		v := Average(pair.values)
// 		if v != pair.average {
// 			t.Error(
// 				"For", pair.values,
// 				"expected", pair.average,
// 				"got", v)
// 		}
// 	}
// }
