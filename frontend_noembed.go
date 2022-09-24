//go:build !embed

package main

import (
	"embed"
	flag "github.com/spf13/pflag"
)

var frontendEmbed embed.FS

var defaultFrontend = "frontend/dist"

func init() {
	flag.String("frontend", defaultFrontend, "Override frontend asset directory.")
}
