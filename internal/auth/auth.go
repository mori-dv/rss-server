package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("no authentication info")
	}
	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authentication info")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of authentication header")
	}
	return vals[1], nil
}
