package http_test

import (
	"fmt"
	"net/http"
	"testing"

	slinkyhttp "github.com/skip-mev/slinky/pkg/http"
	"github.com/stretchr/testify/require"
)

func TestRoundTripperWithHeaders(t *testing.T) {
	expectedHeaderFields := map[string]string{
		"X-Api-Key": "test",
	}

	rt := &customRoundTripper{
		expectedHeaderFields: expectedHeaderFields,
	}

	rtWithHeaders := slinkyhttp.NewRoundTripperWithHeaders(expectedHeaderFields, rt)

	client := &http.Client{
		Transport: rtWithHeaders,
	}

	req, err := http.NewRequest("GET", "http://test.com", nil)
	require.NoError(t, err)

	// Make the request
	_, err = client.Do(req)
	require.NoError(t, err)
}

type customRoundTripper struct {
	expectedHeaderFields map[string]string
}

func (c *customRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range c.expectedHeaderFields {
		if req.Header.Get(k) != v {
			return nil, fmt.Errorf("expected header %s to be %s, got %s", k, v, req.Header.Get(k))
		}
	}
	return &http.Response{}, nil
}
