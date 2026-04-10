package solver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RequestErr struct {
	StatusCode int
	Contents   string
}

func (ue *RequestErr) Error() string {
	return fmt.Sprintf("unexpected status code %d with contents %s", ue.StatusCode, ue.Contents)
}

func (ue *RequestErr) Is(err error) bool {
	if e, ok := errors.AsType[*RequestErr](err); ok {
		return e.Contents == ue.Contents && e.StatusCode == ue.StatusCode
	}
	return false
}

type RemoteSolver struct {
	MathServerURL string
	Client        *http.Client
}

func (rs RemoteSolver) Resolve(ctx context.Context, expression string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		rs.MathServerURL+"?expression="+url.QueryEscape(expression), nil)
	if err != nil {
		return 0, err
	}
	resp, err := rs.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, &RequestErr{
			StatusCode: resp.StatusCode,
			Contents:   string(contents),
		}
	}
	result, err := strconv.ParseFloat(string(contents), 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
