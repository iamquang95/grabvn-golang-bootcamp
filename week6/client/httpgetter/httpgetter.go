package httpgetter

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/avast/retry-go"
)

type HTTPGetter interface {
	Get(url string) ([]byte, error)
}

type HTTPGetterWithRetry struct {
	maxRetry uint
}

func NewHTTPGetterWithRetry(maxRetry uint) HTTPGetter {
	httpGetter := &HTTPGetterWithRetry{
		maxRetry: maxRetry,
	}
	return httpGetter
}

func (httpGetter *HTTPGetterWithRetry) Get(url string) ([]byte, error) {
	var body []byte
	err := retry.Do(
		func() error {
			resp, err := http.Get(url)

			if err != nil {
				return err
			}
			if resp.StatusCode == 500 || resp.StatusCode == 501 {
				return errors.New("Internal server error")
			}

			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)

			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(httpGetter.maxRetry),
	)
	return body, err
}
