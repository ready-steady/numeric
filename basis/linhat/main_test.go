package linhat

import (
	"reflect"
	"testing"
)

func assertEqual(actual, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("got '%v' instead of '%v'", actual, expected)
	}
}

func TestEvaluate(t *testing.T) {
	basis := New(1)

	points := []float64{-1, 0, 0.25, 0.5, 0.75, 1, 2}

	cases := []struct {
		level  uint8
		order  uint32
		values []float64
	}{
		{0, 0, []float64{0, 1, 1.0, 1, 1.0, 1, 0}},
		{1, 0, []float64{0, 1, 0.5, 0, 0.0, 0, 0}},
		{1, 2, []float64{0, 0, 0.0, 0, 0.5, 1, 0}},
		{2, 1, []float64{0, 0, 1.0, 0, 0.0, 0, 0}},
		{2, 3, []float64{0, 0, 0.0, 0, 1.0, 0, 0}},
	}

	values := make([]float64, len(points))

	for i := range cases {
		for j := range values {
			values[j] = basis.Evaluate([]float64{points[j]},
				[]uint8{cases[i].level}, []uint32{cases[i].order})
		}
		assertEqual(values, cases[i].values, t)
	}
}
