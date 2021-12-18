//go:build ignore

package main

import (
	flag "github.com/spf13/pflag"
	"io"
	"log"
	"net/http"
	"os"
)

const OUT = ".cmudict.dict"

func main() {
	var location string
	flag.StringVarP(&location,
		"url",
		"u",
		"https://github.com/cmusphinx/cmudict/raw/master/cmudict.dict",
		"Source URL",
	)

	var output string
	flag.StringVarP(&output, "output", "o", ".cmudict.dict", "Output filename")

	flag.Parse()

	log.Println("Downloading CMUdict from " + location)
	resp, err := http.Get(location)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Server returned \"%s\"", resp.Status)
	}

	log.Println("Saving into " + output)
	out, err := os.Create(output)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(out, resp.Body)

	err = out.Close()
	if err != nil {
		log.Fatalln(err)
	}

}
