package testmain

import (
	"log/slog"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	slog.Info("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	slog.Info("Clean up stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	t.Log("TestFirst uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	t.Log("TestSecond also uses stuff set up in TestMain", testTime)
}
