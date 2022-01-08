//go:generate go run internal/cmudict.go

package main

import (
	"embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal/database"
	"github.com/gabe565/mnemonic-ninja/internal/server"
	"github.com/gabe565/mnemonic-ninja/internal/word"
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

	db, err := database.SetupDatabase()
	if err != nil {
		panic(err)
	}

	err = word.ImportWords(db, cmudict)
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
	router := server.Router(db, contentFs)
	log.Println("Listening on " + *address)
	err = http.ListenAndServe(*address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
