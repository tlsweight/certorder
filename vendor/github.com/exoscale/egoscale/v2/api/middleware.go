package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

type Middleware interface {
	http.RoundTripper
}

// ErrorHandlerMiddleware is an Exoscale API HTTP client middleware that
// returns concrete Go errors according to API response errors.
type ErrorHandlerMiddleware struct {
	next http.RoundTripper
}

func NewAPIErrorHandlerMiddleware(next http.RoundTripper) Middleware {
	if next == nil {
		next = http.DefaultTransport
	}

	return &ErrorHandlerMiddleware{next: next}
}

func (m *ErrorHandlerMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := m.next.RoundTrip(req)
	if err != nil {
		// If the request returned a Go error don't bother analyzing the response
		// body, as there probably won't be any (e.g. connection timeout/refused).
		return resp, err
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		var res struct {
			Message string `json:"message"`
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %s", err)
		}

		if json.Valid(data) {
			if err = json.Unmarshal(data, &res); err != nil {
				return nil, fmt.Errorf("error unmarshaling response: %s", err)
			}
		} else {
			res.Message = string(data)
		}

		switch {
		case resp.StatusCode == http.StatusNotFound:
			return nil, ErrNotFound

		case resp.StatusCode >= 400 && resp.StatusCode < 500:
			return nil, fmt.Errorf("%w: %s", ErrInvalidRequest, res.Message)

		case resp.StatusCode >= 500:
			return nil, fmt.Errorf("%w: %s", ErrAPIError, res.Message)
		}
	}

	return resp, err
}

// TraceMiddleware is a client HTTP middleware that dumps HTTP requests and responses content.
type TraceMiddleware struct {
	next http.RoundTripper
}

func NewTraceMiddleware(next http.RoundTripper) Middleware {
	if next == nil {
		next = http.DefaultTransport
	}

	return &TraceMiddleware{next: next}
}

func (t *TraceMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	if dump, err := httputil.DumpRequest(req, true); err == nil {
		fmt.Fprintf(os.Stderr, ">>> %s\n", dump)
	}

	fmt.Fprintln(os.Stderr, "----------------------------------------------------------------------")

	resp, err := t.next.RoundTrip(req)

	if resp != nil {
		if dump, err := httputil.DumpResponse(resp, true); err == nil {
			fmt.Fprintf(os.Stderr, "<<< %s\n", dump)
		}
	}

	return resp, err
}
