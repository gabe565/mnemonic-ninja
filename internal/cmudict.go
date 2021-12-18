package internal

import (
	_ "embed"
)

//go:embed cmudict.go
var CMUdict string

const CMUDICT_URL = "https://github.com/cmusphinx/cmudict/raw/master/cmudict.dict"
