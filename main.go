//go:generate go run cmudict.go
//go:generate -command npm npm -prefix frontend
//go:generate npm install
//go:generate npm run build

package main

import (
	"embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal"
	flag "github.com/spf13/pflag"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed .cmudict.dict
var cmudict string

//go:embed frontend/dist
var dist embed.FS

func main() {
	var err error

	address := flag.String("address", ":3000", "Override listen address.")
	staticDir := flag.String("static", "", "Override static asset directory. Useful for development. If left empty, embedded assets are used.")
	flag.Parse()

	err = internal.SetupDatabase()
	if err != nil {
		panic(err)
	}

	err = internal.ImportWords(cmudict)
	if err != nil {
		panic(err)
	}

	var contentFs fs.FS
	if *staticDir != "" {
		contentFs = os.DirFS(*staticDir)
	} else {
		contentFs, err = fs.Sub(dist, "frontend/dist")
		if err != nil {
			panic(err)
		}
	}
	router := internal.Router(contentFs)
	log.Println("Listening on " + *address)
	err = http.ListenAndServe(*address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
