package ex1

import "testing"

// the function for testing Average should be named TestAverage and have a single parameter named t of type *testing.T
func TestAverage(t *testing.T) {
	data := []float64{10, 20, 30}
	if average := Average(data); average != 20 {
		// use Errorf to report the error
		t.Errorf("expected average to be 20, got %f", average)
	}
}
