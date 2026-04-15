package ex4

import (
	"errors"
	"testing"
)

func TestCalcWealth(t *testing.T) {
	data := []struct {
		name  string
		gp    int
		sp    int
		bp    int
		total int
		err   error
	}{
		{
			name:  "all_bronze",
			gp:    0,
			sp:    0,
			bp:    10,
			total: 10,
			err:   nil,
		},
		{
			name:  "all_silver",
			gp:    0,
			sp:    10,
			bp:    0,
			total: 100,
			err:   nil,
		},
		{
			name:  "silver_and_bronze",
			gp:    0,
			sp:    10,
			bp:    10,
			total: 110,
			err:   nil,
		},
		{
			name:  "negative_bronze",
			gp:    0,
			sp:    0,
			bp:    -1,
			total: 0,
			err:   ErrNegative,
		},
		{
			name:  "negative_silver",
			gp:    10,
			sp:    -1,
			bp:    10,
			total: 0,
			err:   ErrNegative,
		},
		{
			name:  "negative_gold",
			gp:    -5,
			sp:    3,
			bp:    10,
			total: 0,
			err:   ErrNegative,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			total, err := CalcWealth(d.gp, d.sp, d.bp)
			if !errors.Is(err, d.err) {
				t.Errorf("expected %v, got %v", d.err, err)
			}
			if total != d.total {
				t.Errorf("expected %d, got %d", d.total, total)
			}
		})
	}
}
