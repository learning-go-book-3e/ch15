package oracle_test

import (
	"testing"

	"github.com/learning-go-book-3e/ch15/sample_code/oracle"
)

func TestOracle(t *testing.T) {
	ctx := t.Context()
	ch := oracle.Launch(ctx)
	outCh := make(chan string)
	ch <- oracle.Request{
		Query:    "Is the sky blue?",
		Response: outCh,
	}
	result := <-outCh
	if result != "Is the sky blue? Yes!" {
		t.Error("the oracle is wearing rose-tinted glasses")
	}
}
