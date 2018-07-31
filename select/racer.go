package racer

import (
	"net/http"
	"time"
)

func Racer(first_url, second_url string) (winner string) {
	firstDuration := measureResponseTime(first_url)
	secondDuration := measureResponseTime(second_url)

	if firstDuration < secondDuration {
		return first_url
	}

	return second_url
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
