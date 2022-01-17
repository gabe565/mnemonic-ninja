//go:generate go run internal/cmudict.go

package main

import (
	"embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal/database"
	"github.com/gabe565/mnemonic-ninja/internal/database/seeds"
	"github.com/gabe565/mnemonic-ninja/internal/server"
	flag "github.com/spf13/pflag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

const EnvPrefix = "MNEMONIC_NINJA_"

//go:embed .cmudict.dict
var cmudict string

//go:embed frontend/dist
var dist embed.FS

func main() {
	var err error

	address := flag.String("address", ":3000", "Override listen address.")
	staticDir := flag.String("static", "", "Override static asset directory. Useful for development. If left empty, embedded assets are used.")
	flag.Parse()

	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		optName := strings.ToUpper(f.Name)
		optName = strings.ReplaceAll(optName, "-", "_")
		varName := EnvPrefix + optName
		if val, ok := os.LookupEnv(varName); !f.Changed && ok {
			err = f.Value.Set(val)
			if err != nil {
				log.Fatalln(err)
			}
		}
	})

	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	err = seeds.SeedWords(db, cmudict)
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