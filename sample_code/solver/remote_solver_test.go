package solver

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {
	type info struct {
		expression string
		code       int
		body       string
	}
	var io info
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			// validate the incoming request to make sure the client is
			// sending the expected data
			if req.URL.Query().Get("expression") != io.expression {
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
			rw.WriteHeader(io.code)
			rw.Write([]byte(io.body))
		}))
	defer server.Close()
	rs := RemoteSolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	}
	data := []struct {
		name   string
		io     info
		result float64
		err    error
	}{
		{"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22, nil},
		{"case2", info{"( 2 + 2 ) * 10", http.StatusOK, "40"}, 40, nil},
		{"case3", info{"( 2 + 2 * 10", http.StatusBadRequest,
			"invalid expression: ( 2 + 2 * 10"}, 0,
			&RequestErr{
				StatusCode: http.StatusBadRequest,
				Contents:   "invalid expression: ( 2 + 2 * 10",
			}},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(t.Context(), d.io.expression)
			if result != d.result {
				t.Errorf("io `%f`, got `%f`", d.result, result)
			}
			if !errors.Is(err, d.err) {
				t.Errorf("expected error %v, got %v", d.err, err)
			}
		})
	}
}
