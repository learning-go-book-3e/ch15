package bench

import (
	"fmt"
	"math/rand/v2"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	makeData()
	exitVal := m.Run()
	os.Remove("testdata/data.txt")
	os.Exit(exitVal)
}

// makeData makes our data file for us. Rather than checking in a large file, we recreate it for the test.
// By setting the random seed to the same value every time, we ensure that we generate the same file every time.
// This random seed generates a file that's 64,847 bytes long.
func makeData() {
	file, err := os.Create("testdata/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := rand.New(rand.NewPCG(0, 0))
	for range 10_000 {
		data := makeWord(r, r.IntN(10)+1)
		file.Write(data)
	}
}

func makeWord(r *rand.Rand, wordLen int) []byte {
	out := make([]byte, wordLen+1)
	for i := range wordLen {
		out[i] = 'a' + byte(r.IntN(26))
	}
	out[wordLen] = '\n'
	return out
}

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != 64_847 {
		t.Error("Expected 64,847, got", result)
	}
}

func BenchmarkFileLen1(b *testing.B) {
	for b.Loop() {
		_, err := FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for b.Loop() {
				_, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
