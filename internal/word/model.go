package word

import (
	"database/sql"
	"regexp"
	"strings"
)

type WordModel struct {
	ID      uint           `gorm:"primaryKey"`
	Word    sql.NullString `gorm:"index"`
	Arpabet string         `gorm:"-"`
	Number  sql.NullString `gorm:"index"`
}

func (WordModel) TableName() string {
	return "words"
}

var numberRegex = regexp.MustCompile("[^0-9]")

func New(s string) (*WordModel, error) {
	var err error
	w := &WordModel{}

	// Split word and arpabet
	split := strings.SplitN(s, " ", 2)

	// Remove (#) from duplicate words
	word := strings.SplitN(split[0], "(", 2)[0]
	err = w.Word.Scan(word)
	if err != nil {
		return w, err
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
		return w, err
	}

	return w, nil
}
