package word

import "strings"

type Word struct {
	ID      uint   `json:"-"`
	Word    string `json:"word"`
	Arpabet string `json:"arpabet"`
	Number  string `json:"number"`
}

func New(s string) *Word {
	// Split word and arpabet
	split := strings.SplitN(s, " ", 2)

	// Remove (#) from duplicate words
	split[0] = strings.SplitN(split[0], "(", 2)[0]

	// Remove comments
	split[1] = strings.SplitN(split[1], "#", 2)[0]

	return &Word{
		Word:    split[0],
		Arpabet: split[1],
		Number:  strings.Trim(ArpabetReplacer.Replace(split[1]+" "), " "),
	}
}

func (w Word) String() string {
	return w.Word + "\t\t" + w.Arpabet + "\t\t" + w.Number
}
