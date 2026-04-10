package artifact

import (
	"os"
	"testing"
)

func TestArtifacts(t *testing.T) {
	myDir := t.ArtifactDir()
	t.Log(myDir)
	f, err := os.Create(myDir + "/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("valuable test data"))
}
