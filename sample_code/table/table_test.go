package table

import (
	"errors"
	"testing"
)

func TestDoMath(t *testing.T) {
	result, err := DoMath(2, 2, "+")
	if result != 4 {
		t.Error("Should have been 4, got", result)
	}
	if err != nil {
		t.Error("Should have been nil error, got", err)
	}
	result2, err2 := DoMath(2, 2, "-")
	if result2 != 0 {
		t.Error("Should have been 0, got", result2)
	}
	if err2 != nil {
		t.Error("Should have been nil error, got", err2)
	}
	result3, err3 := DoMath(2, 2, "*")
	if result3 != 4 {
		t.Error("Should have been 4, got", result3)
	}
	if err3 != nil {
		t.Error("Should have been nil error, got", err3)
	}
	result4, err4 := DoMath(2, 2, "/")
	if result4 != 1 {
		t.Error("Should have been 1, got", result4)
	}
	if err4 != nil {
		t.Error("Should have been nil error, got", err4)
	}
}

func TestDoMathTable(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		err      error
	}{
		{"addition", 2, 2, "+", 4, nil},
		{"subtraction", 2, 2, "-", 0, nil},
		{"multiplication", 2, 2, "*", 4, nil},
		{"division", 2, 2, "/", 1, nil},
		{"bad_division", 2, 0, "/", 0, ErrDivZero},
		{"bad_op", 2, 2, "?", 0, &UnknownOpErr{Op: "?"}},
		{"another_mult", 2, 3, "*", 6, nil},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			if !errors.Is(err, d.err) {
				t.Errorf("expected %v, got %v", d.err, err)
			}
		})
	}
}
