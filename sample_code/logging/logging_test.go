package logging

import (
	"fmt"
	"testing"
)

func TestLogging(t *testing.T) {
	fmt.Println("fmt.Println does work, but it's always there and unindented")
	t.Log("This is a simple string message")
	t.Logf("This message has %d %s specifiers in it", 2, "formatting")
	t.Attr("key_for_processing", "value_for_processing with whitespace")
	t.Output().Write([]byte("a final entry"))
}
