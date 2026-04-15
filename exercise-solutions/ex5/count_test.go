package ex5

import (
	"math/rand/v2"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	makeData()
	exitVal := m.Run()
	os.Exit(exitVal)
}

var sampleString string
var numWords int

// makeData makes our data for us.
// This random seed generates a string that's 2226 words long.
func makeData() {
	r := rand.New(rand.NewPCG(0, 0))
	numWords = r.IntN(10_000)
	var sb strings.Builder

	for range numWords {
		sb.Write(makeWord(r, r.IntN(10)+1))
		sb.WriteString(" ")
	}
	sampleString = sb.String()
}

func makeWord(r *rand.Rand, wordLen int) []byte {
	out := make([]byte, wordLen)
	for i := range wordLen {
		out[i] = 'a' + byte(r.IntN(26))
	}
	return out
}

func TestCounters(t *testing.T) {
	resultFields := CountWithFields(sampleString)
	if resultFields != numWords {
		t.Errorf("CountWithFields: expected %d, got %d", numWords, resultFields)
	}

	resultFieldsSeq := CountWithFieldsSeq(sampleString)
	if resultFieldsSeq != numWords {
		t.Errorf("CountWithFieldsSeq: expected %d, got %d", numWords, resultFieldsSeq)
	}

	resultManual := CountManual(sampleString)
	if resultManual != numWords {
		t.Errorf("CountManual: expected %d, got %d", numWords, resultManual)
	}
}

func BenchmarkCountWithFields(b *testing.B) {
	for b.Loop() {
		CountWithFields(sampleString)
	}
}

func BenchmarkCountWithFieldsSeq(b *testing.B) {
	for b.Loop() {
		CountWithFieldsSeq(sampleString)
	}
}

func BenchmarkCountManual(b *testing.B) {
	for b.Loop() {
		CountManual(sampleString)
	}
}
