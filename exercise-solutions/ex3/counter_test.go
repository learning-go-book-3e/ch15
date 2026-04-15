package ex3

import (
	"math/rand/v2"
	"os"
	"strconv"
	"testing"
)

func TestTotalFile(t *testing.T) {
	// create a file to load up
	fileNameGood, totalGood := makeDataGood(t)
	fileNameBad, _ := makeDataBad(t)

	// three cases, one with all good, one with bad data, one with bad filename
	result, err := TotalFile(fileNameGood)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != totalGood {
		t.Errorf("expect %d, got %d", totalGood, result)
	}

	result2, err2 := TotalFile(fileNameBad)
	if err2 == nil {
		t.Errorf("expected error, got nil")
	}
	if result2 != 0 {
		t.Errorf("expect 0, got %d", result2)
	}

	result3, err3 := TotalFile("qwijibo")
	if err3 == nil {
		t.Errorf("expected error, got nil")
	}
	if result3 != 0 {
		t.Errorf("expect 0, got %d", result3)
	}

}

func makeDataGood(t *testing.T) (string, int) {
	t.Helper()
	tempDir := t.TempDir()
	f, err := os.CreateTemp(tempDir, "tempFile")
	if err != nil {
		t.Fatal(err)
	}
	// generate 10 numbers and write them to the file
	total := 0
	for range 10 {
		i := rand.Int()
		total += i
		f.WriteString(strconv.Itoa(i))
		f.WriteString("\n")
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	return f.Name(), total
}

func makeDataBad(t *testing.T) (string, int) {
	t.Helper()
	tempDir := t.TempDir()
	f, err := os.CreateTemp(tempDir, "tempFile")
	if err != nil {
		t.Fatal(err)
	}
	// generate 3 numbers and write them to the file
	total := 0
	for range 3 {
		i := rand.Int()
		total += i
		f.WriteString(strconv.Itoa(i))
		f.WriteString("\n")
	}
	// write a word
	f.WriteString("oops\n")
	// generate 3 more numbers and write to the file
	for range 3 {
		i := rand.Int()
		total += i
		f.WriteString(strconv.Itoa(i))
		f.WriteString("\n")
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	return f.Name(), total
}
