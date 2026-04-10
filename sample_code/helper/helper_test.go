package helper

import (
	"io"
	"os"
	"testing"
)

func TestSomething(t *testing.T) {
	f := openFile(t, "/tmp/does_not_exist.txt")
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty file")
	}
}

func openFile(t *testing.T, name string) io.Reader {
	t.Helper()
	out, err := os.Open(name)
	if err != nil {
		t.Fatal("could not open file:", err)
	}
	t.Cleanup(func() {
		_ = out.Close()
	})
	return out
}
