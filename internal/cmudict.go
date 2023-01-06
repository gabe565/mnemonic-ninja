//go:build ignore

package main

import (
	"compress/gzip"
	flag "github.com/spf13/pflag"
	"io"
	"log"
	"net/http"
	"os"
)

// renovate branch=master
const Url = "https://github.com/cmusphinx/cmudict/raw/697cd8957daa07d2cb89e0cee922de6218afde81/cmudict.dict"

func main() {
	var output string
	flag.StringVarP(&output, "output", "o", ".cmudict.dict.gz", "Output filename")

	flag.Parse()

	log.Println("Downloading CMUdict from " + Url)
	resp, err := http.Get(Url)
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

	gz := gzip.NewWriter(out)

	if _, err := io.Copy(gz, resp.Body); err != nil {
		log.Fatalln(err)
	}

	if err := gz.Close(); err != nil {
		log.Fatalln(err)
	}

	if err := out.Close(); err != nil {
		log.Fatalln(err)
	}

}
