package word

import (
	"regexp"
	"strings"
)

type WordModel struct {
	Word    string `gorm:"index"`
	Arpabet string `gorm:"-"`
	Number  string `gorm:"index"`
	Guess   bool   `gorm:"-"`
}

func (WordModel) TableName() string {
	return "words"
}

var numberRegex = regexp.MustCompile("[^0-9]")

func FromCmudict(line string) *WordModel {
	w := WordModel{}

	// Split word and arpabet
	split := strings.SplitN(line, " ", 2)

	// Remove (#) from duplicate words
	w.Word = strings.SplitN(split[0], "(", 2)[0]

	// Remove comments
	arpabet := strings.SplitN(split[1], "#", 2)[0]
	arpabet = strings.Trim(arpabet, " ")
	w.Arpabet = arpabet

	for _, v := range strings.SplitAfter(arpabet+" ", " ") {
		number := ArpabetReplacer.Replace(v)
		if numberRegex.MatchString(number) {
			continue
		}
		w.Number += number
	}

	return &w
}

func FromString(word string) *WordModel {
	number := LetterReplacer.Replace(word)
	number = numberRegex.ReplaceAllLiteralString(number, "")

	return &WordModel{
		Word:   word,
		Number: number,
		Guess:  true,
	}
}
