package table

import (
	"errors"
	"fmt"
)

var ErrDivZero = errors.New("division by zero")

type UnknownOpErr struct {
	Op string
}

func (ue *UnknownOpErr) Error() string {
	return fmt.Sprintf("unknown operator %s", ue.Op)
}

func (ue *UnknownOpErr) Is(err error) bool {
	if e, ok := errors.AsType[*UnknownOpErr](err); ok && e.Op == ue.Op {
		return true
	}
	return false
}

func DoMath(num1, num2 int, op string) (int, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 + num2, nil
	case "/":
		if num2 == 0 {
			return 0, ErrDivZero
		}
		return num1 / num2, nil
	default:
		return 0, &UnknownOpErr{Op: op}
	}
}
