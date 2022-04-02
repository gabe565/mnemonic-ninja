package models

import (
	"errors"
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"regexp"
	"strings"
)

type Word struct {
	Word    string `gorm:"type:text collate nocase;index"`
	Arpabet string `gorm:"-"`
	Number  string `gorm:"index"`
	Guess   bool   `gorm:"-"`
}

func (Word) TableName() string {
	return "words"
}

var numberRegex = regexp.MustCompile("[^0-9]")

var ErrInvalidArpabet = errors.New("invalid arpabet")

func FromCmudict(line string) (*Word, error) {
	w := Word{}

	// Split word and arpabet
	split := strings.SplitN(line, " ", 2)

	// Remove (#) from duplicate words
	w.Word = strings.SplitN(split[0], "(", 2)[0]

	// Remove comments
	arpabet := strings.SplitN(split[1], "#", 2)[0]
	arpabet = strings.Trim(arpabet, " ")
	w.Arpabet = arpabet
	if w.Arpabet == "" {
		return &w, fmt.Errorf("%v: %w", w.Word, ErrInvalidArpabet)
	}

	for _, v := range strings.SplitAfter(arpabet+" ", " ") {
		number := word.Arpabet.Replace(v)
		if numberRegex.MatchString(number) {
			return &w, fmt.Errorf("%v: %w", w.Word, ErrInvalidArpabet)
		}
		w.Number += number
	}

	return &w, nil
}

func FromString(w string) *Word {
	w = strings.ToLower(w)

	number := word.Letter.Replace(w)
	number = numberRegex.ReplaceAllLiteralString(number, "")

	return &Word{
		Word:   w,
		Number: number,
		Guess:  true,
	}
}
