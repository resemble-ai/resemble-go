package util

import (
	"encoding/json"
	"fmt"
	"log"
)

// ApiError represent API errror data model
type ApiError struct {
	Method     string `json:"-"`
	StatusCode int    `json:"-"`
	Endpoint   string `json:"-"`
	Message    string `json:"message"`
}

func (a *ApiError) Error() string {
	return fmt.Sprintf("%s /%s returned %d %s", a.Method, a.Endpoint, a.StatusCode, a.Message)
}

// NewApiError returns new ApiError instance
func NewApiError(body []byte, endpoint string, statusCode int, method string) error {
	apiError := ApiError{
		Endpoint:   endpoint,
		StatusCode: statusCode,
		Method:     method,
	}
	if err := json.Unmarshal(body, &apiError); err != nil {
		log.Println("can not convert to ApiError:", string(body))
		return fmt.Errorf("%s", body)
	}
	return &apiError
}
