package parallel

import (
	"testing"
)

func TestParallelTable(t *testing.T) {
	data := []struct {
		name   string
		input  int
		output int
	}{
		{"a", 10, 20},
		{"b", 30, 40},
		{"c", 50, 60},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			t.Log(d.input, d.output)
			out := toTest(d.input)
			if out != d.output {
				t.Error("didn't match", out, d.output)
			}
		})
	}
}
