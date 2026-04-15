package ex1

func Average(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	// the formula for an average was wrong
	return total / float64(len(nums))
}
