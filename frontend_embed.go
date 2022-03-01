//go:build embed

package main

//go:generate -command npm npm --prefix frontend
//go:generate npm install
//go:generate npm run build

import "embed"

//go:embed frontend/dist
var frontendEmbed embed.FS

var defaultFrontend string
var frontendHelpExt = " If left empty, embedded assets are used."
