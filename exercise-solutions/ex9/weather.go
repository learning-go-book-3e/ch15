package ex9

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Client communicates with a weather API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

var (
	ErrCityMissing = errors.New("city is required")
)

type RequestFailedErr struct {
	StatusCode int
	Body       string
}

func (be *RequestFailedErr) Error() string {
	return fmt.Sprintf("bad request: status code: %d, message: %s", be.StatusCode, be.Body)
}

func (be *RequestFailedErr) Is(err error) bool {
	if e, ok := errors.AsType[*RequestFailedErr](err); ok {
		return e.StatusCode == be.StatusCode && e.Body == be.Body
	}
	return false
}

type BadResponseErr struct {
	Body string
}

func (be *BadResponseErr) Error() string {
	return fmt.Sprintf("bad response: cannot parse %q", be.Body)
}

func (be *BadResponseErr) Is(err error) bool {
	if e, ok := errors.AsType[*BadResponseErr](err); ok {
		return e.Body == be.Body
	}
	return false
}

// GetTemperature fetches the temperature for a city.
// The API returns the temperature as a plain-text number.
func (c Client) GetTemperature(ctx context.Context, city string) (float64, error) {
	city = strings.TrimSpace(city)
	if city == "" {
		return 0, ErrCityMissing
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		c.BaseURL+"/temperature?city="+city, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, &RequestFailedErr{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}
	}

	temp, err := strconv.ParseFloat(strings.TrimSpace(string(body)), 64)
	if err != nil {
		return 0, &BadResponseErr{Body: string(body)}
	}

	return temp, nil
}
