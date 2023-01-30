package cmudict

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// renovate branch=master
const Url = "https://github.com/cmusphinx/cmudict/raw/697cd8957daa07d2cb89e0cee922de6218afde81/cmudict.dict"

var ErrUnexpectedResponse = errors.New("unexpected response")

func Download() (io.ReadCloser, error) {
	log.Println("Downloading CMUdict from " + Url)
	resp, err := http.Get(Url)
	if err != nil {
		return resp.Body, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp.Body, fmt.Errorf("%w: %d", ErrUnexpectedResponse, resp.StatusCode)
	}

	return resp.Body, nil
}
