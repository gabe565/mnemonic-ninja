package word

import "strings"

type Word struct {
	ID      uint   `json:"-" gorm:"primaryKey"`
	Word    *string `json:"word,omitempty" gorm:"index"`
	Arpabet string `json:"arpabet"`
	Number  *string `json:"number,omitempty" gorm:"index"`
}

func New(s string) *Word {
	// Split word and arpabet
	split := strings.SplitN(s, " ", 2)

	// Remove (#) from duplicate words
	word := strings.SplitN(split[0], "(", 2)[0]

	// Remove comments
	number := strings.SplitN(split[1], "#", 2)[0]
	number = strings.Trim(ArpabetReplacer.Replace(split[1]+" "), " ")

	return &Word{
		Word:    &word,
		Arpabet: split[1],
		Number:  &number,
	}
}

func (w Word) String() string {
	return (*w.Word) + "\t\t" + w.Arpabet + "\t\t" + (*w.Number)
}
