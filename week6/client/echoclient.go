package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/avast/retry-go"
)

func main() {
	url := "http://127.0.0.1:8080/test"

	body, err := getHTTPRequestWithRetry(url, 2)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))

}

func getHTTPRequestWithRetry(url string, maxRetry uint) ([]byte, error) {
	var body []byte
	err := retry.Do(
		func() error {
			resp, err := http.Get(url)

			if err != nil {
				return err
			}
			if resp.StatusCode == 500 || resp.StatusCode == 501 {
				fmt.Println("Got internal server error")
				return errors.New("Internal server error")
			}

			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)

			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(maxRetry),
	)
	return body, err
}
