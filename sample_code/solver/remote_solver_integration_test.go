//go:build integration

package solver

import (
	"errors"
	"net/http"
	"testing"
)

func TestRemoteSolver_ResolveIntegration(t *testing.T) {
	rs := RemoteSolver{
		MathServerURL: "http://localhost:8080",
		Client:        http.DefaultClient,
	}
	data := []struct {
		name       string
		expression string
		result     float64
		err        error
	}{
		{"case1", "2 + 2 * 10", 22, nil},
		{"case2", "( 2 + 2 ) * 10", 40, nil},
		{"case3", "( 2 + 2 * 10", 0, &RequestErr{
			StatusCode: http.StatusBadRequest,
			Contents:   "invalid expression: ( 2 + 2 * 10",
		}},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := rs.Resolve(t.Context(), d.expression)
			if result != d.result {
				t.Errorf("expected `%f`, got `%f`", d.result, result)
			}
			if !errors.Is(err, d.err) {
				t.Errorf("expected error %v, got %v", d.err, err)
			}
		})
	}
}
