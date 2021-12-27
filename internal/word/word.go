package word

import (
	"regexp"
	"strings"
)

type Word struct {
	ID      uint   `json:"-" gorm:"primaryKey"`
	Word    *string `json:"word,omitempty" gorm:"index"`
	Arpabet string `json:"arpabet"`
	Number  *string `json:"number,omitempty" gorm:"index"`
}

var numberRegex = regexp.MustCompile("[^0-9]")

func New(s string) *Word {
	// Split word and arpabet
	split := strings.SplitN(s, " ", 2)

	// Remove (#) from duplicate words
	word := strings.SplitN(split[0], "(", 2)[0]

	// Remove comments
	arpabet := strings.SplitN(split[1], "#", 2)[0]
	arpabet = strings.Trim(arpabet, " ")

	var numbers string
	for _, v := range strings.SplitAfter(arpabet+" ", " ") {
		number := ArpabetReplacer.Replace(v)
		if numberRegex.MatchString(number) {
			continue
		}
		numbers += number
	}

	return &Word{
		Word:    &word,
		Arpabet: arpabet,
		Number:  &numbers,
	}
}

func (w Word) String() string {
	return (*w.Word) + "\t\t" + w.Arpabet + "\t\t" + (*w.Number)
}
