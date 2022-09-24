//go:build !embed

package main

import "embed"

var frontendEmbed embed.FS

var defaultFrontend = "frontend/dist"
var frontendHelpExt string
