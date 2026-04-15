package ex8

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID    string
	Name  string
	Price decimal.Decimal
}

// DataStore has many methods, but ReportBuilder only uses GetProducts.
type DataStore interface {
	GetProducts(category string) ([]Product, error)
	GetOrders(userID string) ([]string, error)
	GetReviews(productID string) ([]string, error)
	SaveReport(name string, data []byte) error
	DeleteReport(name string) error
}

type ReportBuilder struct {
	Store DataStore
}

// PriceSummary returns a formatted string listing each product
// and its price for the given category.
func (rb ReportBuilder) PriceSummary(category string) (string, error) {
	products, err := rb.Store.GetProducts(category)
	if err != nil {
		return "", err
	}
	if len(products) == 0 {
		return "no products found", nil
	}
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%s: $%v\n", p.Name, p.Price)
	}
	return result, nil
}
