package solver

import (
	"context"
	"errors"
	"fmt"
	"io"
)

var (
	ErrNoExpression = errors.New("no expression to read")
)

type InvalidExpressionErr struct {
	Expression string
}

func (ie *InvalidExpressionErr) Error() string {
	return fmt.Sprintf("invalid expression: %s", ie.Expression)
}

func (ie *InvalidExpressionErr) Is(err error) bool {
	if e, ok := errors.AsType[*InvalidExpressionErr](err); ok {
		return e.Expression == ie.Expression
	}
	return false
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}

type Processor struct {
	Solver MathSolver
}

func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
	curExpression, err := readToNewLine(r)
	if err != nil {
		return 0, err
	}
	if len(curExpression) == 0 {
		return 0, ErrNoExpression
	}
	answer, err := p.Solver.Resolve(ctx, curExpression)
	return answer, err
}

func readToNewLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}
