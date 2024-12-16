package httpClient

import (
	"net/http"
	"time"
)

func NewHttpConnect() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost:     1024,
			MaxIdleConnsPerHost: 1024,
		},
		Timeout: 15 * time.Second,
	}
}
