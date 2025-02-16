package storage

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap" // Assuming you're using zap for logging
)

// newClient creates a new HTTP client with retry logic.
func newClient(accessTokenFunc func() (string, error), retryCount int, retryDelay time.Duration, logger *zap.Logger) *http.Client {
	client := &http.Client{
		Transport: &retryTransport{
			transport: &httpTransport{
				accessTokenFunc: accessTokenFunc,
			},
			retryCount: retryCount,
			retryDelay: retryDelay,
			logger:     logger,
		},
	}

	return client
}

type retryTransport struct {
	transport  http.RoundTripper
	retryCount int
	retryDelay time.Duration
	logger     *zap.Logger
}

// RoundTripWithRetry adds retry logic to the RoundTrip method.
func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var lastErr error

	for i := 0; i < t.retryCount; i++ {
		if i > 0 {
			t.logger.Info("Retrying request", zap.Int("attempt", i+1), zap.String("url", req.URL.String()))
			time.Sleep(t.retryDelay)
		}

		resp, err := t.transport.RoundTrip(req)
		if err == nil {
			// Check status code for retrying
			if resp.StatusCode >= 500 || resp.StatusCode == http.StatusTooManyRequests { // Example retryable status codes
				bodyBytes, _ := io.ReadAll(resp.Body)
				resp.Body.Close() // Important: Close the previous response body
				lastErr = fmt.Errorf("request failed with status: %s, body: %s", resp.Status, string(bodyBytes))
				t.logger.Warn("Request failed with retryable status", zap.Error(lastErr), zap.Int("status_code", resp.StatusCode), zap.String("url", req.URL.String()))
				continue // Retry
			}
			return resp, nil // Success
		}

		lastErr = err
		t.logger.Error("Request attempt failed", zap.Error(lastErr), zap.String("url", req.URL.String()))

	}

	return nil, fmt.Errorf("request failed after %d attempts: %v", t.retryCount, lastErr)
}

type httpTransport struct {
	accessTokenFunc func() (string, error)
}

// RoundTrip is an implementation of the http.RoundTripper interface that adds the Authorization header to requests.
func (t *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.accessTokenFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	return http.DefaultTransport.RoundTrip(req)
}
