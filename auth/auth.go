package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts api key from headers of requests
func GetAPIKey(headers *http.Header) (string, error) {
	val := headers.Get("Authorization");
	if val == "" {
		return "", errors.New("authorization key is required")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("error: Malformed header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("error: Malformed first part of header")
	}

	return vals[1], nil
}