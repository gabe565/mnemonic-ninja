package models

import (
	"github.com/gabe565/mnemonic-ninja/internal/database/models"
)

type ConversionEntry struct {
	Query string `json:"query"`
	Count int64  `json:"count"`
	Guess bool   `json:"guess,omitempty"`

	Words  []*models.Word `json:"-"`
	Result []string       `json:"result"`
}
