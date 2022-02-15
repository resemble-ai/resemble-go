package api

import (
	"context"
	"fmt"
	"net/http"
)

const baseUrl = "https://app.resemble.ai/api/v2"

// Operation Api operation interface method collections
type Operation interface {
	// Post send Post Request by given payload
	Post(ctx context.Context, path string, data interface{}) (*http.Response, error)

	// Put send Put request by given payload
	Put(ctx context.Context, path string, data interface{}) (*http.Response, error)

	// Patch send Patch request by given payload
	Patch(ctx context.Context, path string, data interface{}) (*http.Response, error)

	// Delete send Delete request by given payload
	Delete(ctx context.Context, path string, data interface{}) (*http.Response, error)

	// Get send Get request by given payload
	Get(ctx context.Context, path string) (*http.Response, error)

	// Stream send Stream request by given payload
	Stream(ctx context.Context, syncServer string, data interface{}) (*http.Response, error)
}

// NewClient returns a new Api instance
func NewClient(token string) Operation {
	return &client{
		baseUrl: baseUrl,
		token:   token,
	}
}

type client struct {
	baseUrl string
	token   string
}

// Post implements Operation.Post method
func (c *client) Post(ctx context.Context, path string, data interface{}) (*http.Response, error) {
	uri := formatUrl(c.baseUrl, path)
	return requestJson(ctx, uri, http.MethodPost, buildHeaders(c.token, false), data)
}

// Put implements Operation.Put method
func (c *client) Put(ctx context.Context, path string, data interface{}) (*http.Response, error) {
	uri := formatUrl(c.baseUrl, path)
	return requestJson(ctx, uri, http.MethodPut, buildHeaders(c.token, false), data)
}

// Patch implements Operation.Patch method
func (c *client) Patch(ctx context.Context, path string, data interface{}) (*http.Response, error) {
	uri := formatUrl(c.baseUrl, path)
	return requestJson(ctx, uri, http.MethodPatch, buildHeaders(c.token, false), data)
}

// Delete implements Operation.Delete method
func (c *client) Delete(ctx context.Context, path string, data interface{}) (*http.Response, error) {
	uri := formatUrl(c.baseUrl, path)
	return requestJson(ctx, uri, http.MethodDelete, buildHeaders(c.token, false), data)
}

// Get implements Operation.Get method
func (c *client) Get(ctx context.Context, path string) (*http.Response, error) {
	uri := formatUrl(c.baseUrl, path)
	return requestJson(ctx, uri, http.MethodGet, buildHeaders(c.token, false), nil)
}

// Stream implements Operation.Stream method
func (c *client) Stream(ctx context.Context, syncServer string, data interface{}) (*http.Response, error) {
	uri := formatUrl(syncServer, "stream")
	return requestJson(ctx, uri, http.MethodPost, buildHeaders(c.token, true), data)
}

func formatUrl(baseUri, path string) string {
	return fmt.Sprintf("%s/%s", baseUri, path)
}

func buildHeaders(token string, syncServer bool) map[string]string {
	headers := map[string]string{}
	if syncServer {
		headers["x-access-token"] = token
	} else {
		headers["Authorization"] = fmt.Sprintf("Token token=%s", token)
	}

	return headers
}
