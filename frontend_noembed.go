//go:build !embed

package main

import "embed"

var frontendEmbed embed.FS

var defaultFrontend = "frontend"
var frontendHelpExt string
