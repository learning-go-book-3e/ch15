package ex9

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTemperature(t *testing.T) {
	type serverResponse struct {
		code int
		body string
	}

	type info struct {
		city     string
		response serverResponse
	}

	// current holds the response the fake server should return
	// for the current test case.
	var curInfo info

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			city := r.URL.Query().Get("city")
			// validate the input!
			if city == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("missing city parameter"))
				return
			}
			// validate the input!
			if curInfo.city != city {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("incorrect city: %q", city)))
				return
			}
			w.WriteHeader(curInfo.response.code)
			w.Write([]byte(curInfo.response.body))
		}))
	t.Cleanup(func() { server.Close() })

	client := Client{
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}

	data := []struct {
		name     string
		info     info
		expected float64
		err      error
	}{
		{
			name: "successful response",
			info: info{
				city:     "seattle",
				response: serverResponse{code: http.StatusOK, body: "72.5"},
			},
			expected: 72.5,
			err:      nil,
		},
		{
			name: "server error",
			info: info{
				city:     "atlantis",
				response: serverResponse{http.StatusNotFound, "city not found"},
			},
			expected: 0,
			err: &RequestFailedErr{
				StatusCode: http.StatusNotFound,
				Body:       "city not found",
			},
		},
		{
			name: "unparseable body",
			info: info{
				city:     "tokyo",
				response: serverResponse{http.StatusOK, "not a number"},
			},
			expected: 0,
			err:      &BadResponseErr{Body: "not a number"},
		},
		{
			name: "empty city",
			info: info{
				city: "",
				// shouldn't get to server
				response: serverResponse{},
			},
			expected: 0,
			err:      ErrCityMissing,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			curInfo = d.info

			result, err := client.GetTemperature(context.Background(), d.info.city)
			if !errors.Is(err, d.err) {
				t.Errorf("expected %v, got %v", d.err, err)
			}
			if math.Abs(result-d.expected) > 0.001 {
				t.Errorf("expected %.2f, got %.2f", d.expected, result)
			}
		})
	}
}
