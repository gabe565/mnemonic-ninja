//go:generate go run internal/cmudict.go

package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal/config"
	"github.com/gabe565/mnemonic-ninja/internal/database"
	"github.com/gabe565/mnemonic-ninja/internal/database/seeds"
	"github.com/gabe565/mnemonic-ninja/internal/server"
	"github.com/gabe565/mnemonic-ninja/internal/telemetry"
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
	var err error

	if err = config.Parse(); err != nil {
		panic(err)
	}

	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	err = seeds.SeedWords(db, cmudictGz)
	if err != nil {
		panic(err)
	}

	var frontendFs fs.FS
	frontendDir := viper.GetString("frontend")
	if frontendDir != "" {
		frontendFs = os.DirFS(frontendDir)
	} else {
		frontendFs, err = fs.Sub(frontendEmbed, "frontend/dist")
		if err != nil {
			panic(err)
		}
	}

	tp, err := telemetry.Setup()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = tp.Shutdown(context.Background())
	}()

	router := server.Router(db, frontendFs, tp)
	address := viper.GetString("address")
	log.Println("Listening on " + address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting")
}
