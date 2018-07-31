package racer

import (
	"net/http"
)

func Racer(first_url, second_url string) (winner string) {
	select {
	case <-ping(first_url):
		return first_url
	case <-ping(second_url):
		return second_url
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
