//go:build ignore

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const CMUDICT_FILE = ".cmudict.dict"
const CMUDICT_URL = "https://github.com/cmusphinx/cmudict/raw/master/cmudict.dict"

func main() {
	log.Println("Downloading CMUdict")
	resp, err := http.Get(CMUDICT_URL)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(CMUDICT_FILE); err == nil {
		log.Println("Removing existing file")
		err = os.Remove(CMUDICT_FILE)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Saving to " + CMUDICT_FILE)
	out, err := os.Create(CMUDICT_FILE)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)
}
