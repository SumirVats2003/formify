package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSON[T any](payload T, r *http.Request) (T, error) {
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
