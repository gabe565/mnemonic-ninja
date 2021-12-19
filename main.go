//go:generate go run internal/download/cmudict.go
//go:generate npm run build

package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal"
	flag "github.com/spf13/pflag"
	"io/fs"
	"net/http"
	"os"
)

//go:embed .cmudict.dict
var cmudict string

//go:embed dist
var dist embed.FS

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

	router := internal.Router(db, dist)
	err = http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
