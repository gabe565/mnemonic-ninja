//go:generate go run internal/download/cmudict.go
package main

import (
	_ "embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal"
	flag "github.com/spf13/pflag"
	"net/http"
	"os"
)

//go:embed .cmudict.dict
var cmudict string

func main() {
	var err error

	var address string
	flag.StringVar(&address, "address", ":3000", "Override listen address.")

	flag.Parse()

	db, err := internal.SetupDatabase()
	if err != nil {
		panic(err)
	}

	err = internal.ImportWords(db, cmudict)
	if err != nil {
		panic(err)
	}

	router := internal.Router(db)
	err = http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
