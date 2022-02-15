package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// sending request by given url, method, header and data
func send(ctx context.Context, url, method string, headers map[string]string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// sending json request by given payload
func requestJson(ctx context.Context, url, method string, headers map[string]string, data interface{}) (*http.Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"

	var (
		payload []byte
		err     error
	)
	if data != nil {
		payload, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	return send(ctx, url, method, headers, payload)
}
