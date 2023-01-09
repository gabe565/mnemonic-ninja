package main

import (
	"github.com/gabe565/mnemonic-ninja/internal/database"
	"github.com/gabe565/mnemonic-ninja/internal/database/seeds"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeeds(t *testing.T) {
	db, err := database.Setup()
	if !assert.NoError(t, err) {
		return
	}

	if err := seeds.SeedWords(db, cmudictGz); !assert.NoError(t, err) {
		return
	}
}
