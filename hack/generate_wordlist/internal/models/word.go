package models

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"

	"github.com/gabe565/mnemonic-ninja/hack/generate_wordlist/internal/word"
)

type Word struct {
	Word    string
	Arpabet string
	Number  string
	Guess   bool
}

func (Word) TableName() string {
	return "words"
}

var numberRegex = regexp.MustCompile("[^0-9]")

var ErrInvalidArpabet = errors.New("invalid arpabet")

func FromCmudict(line []byte) (*Word, error) {
	w := Word{}

	// Split word and arpabet
	split := bytes.SplitN(line, []byte{' '}, 2)

	// Remove (#) from duplicate words
	w.Word = string(truncateAtByte(split[0], '('))

	// Remove comments
	arpabet := bytes.TrimSpace(truncateAtByte(split[1], '#'))
	w.Arpabet = string(arpabet)
	if w.Arpabet == "" {
		return &w, fmt.Errorf("%v: %w", w.Word, ErrInvalidArpabet)
	}

	splitArpabet := bytes.SplitAfter(arpabet, []byte{' '})
	for k, v := range splitArpabet {
		str := string(v)
		if k == len(splitArpabet)-1 {
			str += " "
		}
		number := word.Arpabet.Replace(str)
		if numberRegex.MatchString(number) {
			return &w, fmt.Errorf("%v: %w", w.Word, ErrInvalidArpabet)
		}
		w.Number += number
	}

	return &w, nil
}

func truncateAtByte(b []byte, c byte) []byte {
	i := bytes.IndexByte(b, c)
	if i == -1 {
		return b
	}
	return b[:i]
}
