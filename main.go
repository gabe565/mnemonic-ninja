//go:generate go run cmudict.go
//go:generate npm run build

package main

import (
	"embed"
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

	address := flag.String("address", ":3000", "Override listen address.")
	staticDir := flag.String("static", "", "Override static asset directory. Useful for development. If left empty, embedded assets are used.")
	flag.Parse()

	db, err := internal.SetupDatabase()
	if err != nil {
		panic(err)
	}

	err = internal.ImportWords(db, cmudict)
	if err != nil {
		panic(err)
	}

	var contentFs fs.FS
	if *staticDir != "" {
		contentFs = os.DirFS(*staticDir)
	} else {
		contentFs, err = fs.Sub(dist, "dist")
		if err != nil {
			panic(err)
		}
	}
	router := internal.Router(db, contentFs)
	err = http.ListenAndServe(*address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
