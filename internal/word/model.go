package word

import (
	"database/sql"
	"regexp"
	"strings"
)

type WordModel struct {
	Word    sql.NullString `gorm:"index"`
	Arpabet string         `gorm:"-"`
	Number  sql.NullString `gorm:"index"`
	Guess   bool           `gorm:"-"`
}

func (WordModel) TableName() string {
	return "words"
}

var numberRegex = regexp.MustCompile("[^0-9]")

func FromCmudict(line string) (*WordModel, error) {
	var err error
	w := WordModel{}

	// Split word and arpabet
	split := strings.SplitN(line, " ", 2)

	// Remove (#) from duplicate words
	word := strings.SplitN(split[0], "(", 2)[0]
	err = w.Word.Scan(word)
	if err != nil {
		return &w, err
	}

	// Remove comments
	arpabet := strings.SplitN(split[1], "#", 2)[0]
	arpabet = strings.Trim(arpabet, " ")
	w.Arpabet = arpabet

	var numbers string
	for _, v := range strings.SplitAfter(arpabet+" ", " ") {
		number := ArpabetReplacer.Replace(v)
		if numberRegex.MatchString(number) {
			continue
		}
		numbers += number
	}
	err = w.Number.Scan(numbers)
	if err != nil {
		return &w, err
	}

	return &w, nil
}

func FromString(word string) (*WordModel, error) {
	var err error
	w := WordModel{Guess: true}

	err = w.Word.Scan(word)
	if err != nil {
		return &w, err
	}

	num := LetterReplacer.Replace(word)
	num = numberRegex.ReplaceAllLiteralString(num, "")

	err = w.Number.Scan(num)
	return &w, err
}
