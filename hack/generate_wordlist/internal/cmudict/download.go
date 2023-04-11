package cmudict

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// renovate branch=master
const Url = "https://github.com/cmusphinx/cmudict/raw/7cd8fb5b5a18058688f413e92282eb18815f1956/cmudict.dict"

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
