package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extract an API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey value
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header not found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("authorization header is invalid")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("authorization header is invalid")
	}
	return vals[1], nil
}
