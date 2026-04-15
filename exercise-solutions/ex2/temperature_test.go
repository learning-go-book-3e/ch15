package ex2

import (
	"errors"
	"math"
	"testing"
)

func TestConvert(t *testing.T) {
	data := []struct {
		name   string
		from   string
		to     string
		value  float64
		result float64
		err    error
	}{
		// you should test more than a single number being converted,
		// just to be sure!
		{
			name:   "C to C",
			from:   "C",
			to:     "C",
			value:  100,
			result: 100,
			err:    nil,
		},
		{
			name:   "C to F",
			from:   "C",
			to:     "F",
			value:  100,
			result: 212,
			err:    nil,
		},
		{
			name:   "C to K",
			from:   "C",
			to:     "K",
			value:  100,
			result: 373.15,
			err:    nil,
		},
		{
			name:   "F to C",
			from:   "F",
			to:     "C",
			value:  212,
			result: 100,
			err:    nil,
		},
		{
			name:   "F to F",
			from:   "F",
			to:     "F",
			value:  212,
			result: 212,
			err:    nil,
		},
		{
			name:   "F to K",
			from:   "F",
			to:     "K",
			value:  212,
			result: 373.15,
			err:    nil,
		},
		{
			name:   "K to C",
			from:   "K",
			to:     "C",
			value:  373.15,
			result: 100,
			err:    nil,
		},
		{
			name:   "K to F",
			from:   "K",
			to:     "F",
			value:  373.15,
			result: 212,
			err:    nil,
		},
		{
			name:   "K to K",
			from:   "K",
			to:     "K",
			value:  373.15,
			result: 373.15,
			err:    nil,
		},
		// make sure to test error conditions, too!
		{
			name:  "X to F",
			from:  "X",
			to:    "F",
			value: 12345,
			err:   &UnsupportedUnitError{Name: "X"},
		},
		{
			name:  "K to X",
			from:  "K",
			to:    "X",
			value: 373.15,
			err:   &UnsupportedUnitError{Name: "X"},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := Convert(d.value, d.from, d.to)
			// test errors using Is, not using strings messages!
			if !errors.Is(err, d.err) {
				t.Errorf("expected %v, got %v", d.err, err)
			}

			// Use a small tolerance for floating-point comparison
			if math.Abs(result-d.result) > 0.001 {
				t.Errorf("expected %.4f, got %.4f", d.result, result)
			}
		})
	}
}
