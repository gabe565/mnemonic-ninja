//go:build embed

package main

//go:generate -command npm npm --prefix frontend
//go:generate npm install
//go:generate npm run build

import (
	"embed"
	flag "github.com/spf13/pflag"
)

//go:embed frontend/dist
var frontendEmbed embed.FS

var defaultFrontend string

func init() {
	flag.String("frontend", defaultFrontend, "Override frontend asset directory. If left empty, embedded assets are used.")
}
