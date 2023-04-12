package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gabe565/mnemonic-ninja/hack/generate_wordlist/internal/cmudict"
)

//go:generate go run $GOFILE

func main() {
	var err error

	startTime := time.Now()

	wordlist, err := cmudict.Download()
	if err != nil {
		log.Panic(err)
	}

	dstPath := filepath.Join("..", "..", "src", "assets", "wordlist.csv")
	if err := os.MkdirAll(filepath.Dir(dstPath), 0o777); err != nil {
		log.Panic(err)
	}

	f, err := os.Create(dstPath)
	if err != nil {
		log.Panic(err)
	}

	count, err := cmudict.Csv(f, wordlist)
	if err != nil {
		log.Panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}

	timeTaken := time.Since(startTime)
	log.Printf("Downloaded %d words in %s\n", count, timeTaken)
}
