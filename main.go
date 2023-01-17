//go:generate go run internal/cmudict.go

package main

import (
	"context"
	_ "embed"
	"errors"
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
	"os/signal"
	"syscall"
	"time"
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

	address := viper.GetString("address")
	server := &http.Server{
		Addr:    address,
		Handler: server.Router(db, frontendFs),
	}

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 60 seconds
		ctx, cancelTimeout := context.WithTimeout(ctx, 60*time.Second)
		defer func() {
			cancelTimeout()
		}()

		// Trigger graceful shutdown
		log.Println("Performing graceful shutdown...")
		if err := server.Shutdown(ctx); err != nil {
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				log.Println("Graceful shutdown timed out")
			} else {
				log.Println(err)
			}
		}
		cancel()
	}()

	log.Println("Listening on " + address)
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-ctx.Done()
	log.Println("Exiting")
}
