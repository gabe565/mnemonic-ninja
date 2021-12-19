//go:generate go run internal/download/cmudict.go
package main

import (
	_ "embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal"
	"net/http"
	"os"
)

//go:embed .cmudict.dict
var cmudict string

func main() {
	var err error

	db, err := internal.SetupDatabase()
	if err != nil {
		panic(err)
	}

	err = internal.ImportWords(db, cmudict)
	if err != nil {
		panic(err)
	}

	router := internal.Router(db)
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
