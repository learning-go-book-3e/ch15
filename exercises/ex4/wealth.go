package ex4

import (
	"errors"
)

var (
	ErrNegative = errors.New("negative number of coins")
)

// CalcWealth returns your player's wealth in bronze pieces.
// 10 bronze pieces == 1 silver piece
// 10 gold pieces == 1 silver piece
// A value for any kind of coin returns an error.
func CalcWealth(gp int, sp int, bp int) (int, error) {
	total := 0
	if bp < 0 {
		return 0, ErrNegative
	}
	total += bp
	if sp < 0 {
		return 0, ErrNegative
	}
	total += 10 * sp
	if gp < 0 {
		return 0, ErrNegative
	}
	total += 10 * gp
	return total, nil
}
