//go:generate go run internal/cmudict.go

package main

import (
	_ "embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal/config"
	"github.com/gabe565/mnemonic-ninja/internal/database"
	"github.com/gabe565/mnemonic-ninja/internal/database/seeds"
	"github.com/gabe565/mnemonic-ninja/internal/server"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed .cmudict.dict.gz
var cmudictGz []byte

func init() {
	flag.String("address", ":3000", "Override listen address.")
}

func main() {
	if err := config.Parse(); err != nil {
		panic(err)
	}

	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	go func() {
		if err := seeds.SeedWords(db, cmudictGz); err != nil {
			panic(err)
		}
	}()

	var frontendFs fs.FS
	if frontendDir := viper.GetString("frontend"); frontendDir != "" {
		frontendFs = os.DirFS(frontendDir)
	} else {
		frontendFs, err = fs.Sub(frontendEmbed, "frontend/dist")
		if err != nil {
			panic(err)
		}
	}

	router := server.Router(db, frontendFs)
	address := viper.GetString("address")
	log.Println("Listening on " + address)
	if err := http.ListenAndServe(address, router); err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
