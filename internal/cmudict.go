package internal

import (
	"bufio"
	"log"
	"net/http"
)

const CMUDICT_URL = "https://github.com/cmusphinx/cmudict/raw/master/cmudict.dict"

func GetCmudict() (*bufio.Scanner, error) {
	log.Println("Downloading dictionary")
	response, err := http.Get(CMUDICT_URL)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(response.Body), nil
}
