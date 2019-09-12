package main

import (
	"errors"
	"fmt"

	"grab/week6/client/httpgetter"

	"github.com/myteksi/hystrix-go/hystrix"
)

const (
	url          = "http://127.0.0.1:8080/iamquang95"
	maxRetry     = 2
	maxGetRepeat = 100
)

func main() {
	initCircuitBreaker()
	httpGetter := httpgetter.NewHTTPGetterWithRetry(maxRetry)
	echoCircutBreaker(url, httpGetter)
}

func echoCircutBreaker(url string, httpGetter httpgetter.HTTPGetter) {
	for i := 1; i <= maxGetRepeat; i++ {
		j := i
		hystrix.Do("echo", func() error {
			fmt.Printf("Get echo #%d\n", j)

			body, err := httpGetter.Get(url)

			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println(string(body))
			return err
		}, func(err error) error {
			fmt.Printf("Get echo #%d: fallback, service is down\n", j)
			return errors.New("Fallback")
		})
	}
}

func initCircuitBreaker() {
	hystrix.ConfigureCommand("echo", hystrix.CommandConfig{
		Timeout:                     1000,
		MaxConcurrentRequests:       1,
		ErrorPercentThreshold:       10,
		QueueSizeRejectionThreshold: 100,
		SleepWindow:                 10,
	})
}
