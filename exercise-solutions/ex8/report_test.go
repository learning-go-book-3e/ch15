package ex8

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
)

// PriceSummaryStub embeds DataStore so it satisfies the full interface,
// but only implements GetProducts. Calling any other method will panic.
type PriceSummaryStub struct {
	DataStore
}

var errTestDBTimeout = errors.New("database timeout")

func (ps PriceSummaryStub) GetProducts(category string) ([]Product, error) {
	switch category {
	case "electronics":
		return []Product{
			{ID: "1", Name: "Keyboard", Price: decimal.New(4999, -2)},
			{ID: "2", Name: "Mouse", Price: decimal.New(2999, -2)},
		}, nil
	case "empty":
		return []Product{}, nil
	case "fail":
		return nil, errTestDBTimeout
	default:
		return nil, nil
	}
}

func TestPriceSummary(t *testing.T) {
	data := []struct {
		name     string
		category string
		expected string
		err      error
	}{
		{
			name:     "two products",
			category: "electronics",
			expected: "Keyboard: $49.99\nMouse: $29.99\n",
			err:      nil,
		},
		{
			name:     "no products",
			category: "empty",
			expected: "no products found",
			err:      nil,
		},
		{
			name:     "store error",
			category: "fail",
			expected: "",
			err:      errTestDBTimeout,
		},
	}

	rb := ReportBuilder{Store: PriceSummaryStub{}}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := rb.PriceSummary(d.category)
			if !errors.Is(err, d.err) {
				t.Errorf("expected %v, got %v", d.err, err)
			}
			if result != d.expected {
				t.Errorf("expected %q, got %q", d.expected, result)
			}
		})
	}
}
