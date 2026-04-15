package ex5

import (
	"strings"
	"unicode"
)

// CountWithFields splits on whitespace using strings.Fields
// and returns the number of resulting elements.
func CountWithFields(s string) int {
	return len(strings.Fields(s))
}

// CountWithFieldsSeq splits on whitespace using strings.FieldsSeq
// and returns the number of resulting elements.
func CountWithFieldsSeq(s string) int {
	total := 0
	for range strings.FieldsSeq(s) {
		total++
	}
	return total
}

// CountManual iterates through the string rune by rune,
// counting transitions from whitespace to non-whitespace.
func CountManual(s string) int {
	count := 0
	inWord := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}
	return count
}
