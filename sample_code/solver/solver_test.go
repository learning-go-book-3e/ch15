package solver

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type mathSolverStub struct{}

func (ms mathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, &InvalidExpressionErr{Expression: "( 2 + 2 * 10"}
	}
	return 0, nil
}

func TestProcessor_ProcessExpressions(t *testing.T) {
	p := Processor{mathSolverStub{}}
	in := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10`)
	data := []struct {
		name string
		val  float64
		err  error
	}{
		{name: "case1", val: 22, err: nil},
		{name: "case2", val: 40, err: nil},
		{name: "case3", val: 0, err: &InvalidExpressionErr{Expression: "( 2 + 2 * 10"}},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := p.ProcessExpression(t.Context(), in)
			if !errors.Is(err, d.err) {
				t.Errorf("expected error %v, got %v", d.err, err)
			}
			if result != d.val {
				t.Errorf("Expected result %f, got %f", d.val, result)
			}
		})
	}
}
