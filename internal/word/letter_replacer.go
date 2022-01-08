package word

import "strings"

var LetterReplacer = strings.NewReplacer(
	"th", "1",
	"ch", "6",
	"sh", "6",
	"ck", "7",
	"ph", "8",
	"s", "0",
	"z", "0",
	"t", "1",
	"d", "1",
	"n", "2",
	"m", "3",
	"r", "4",
	"l", "5",
	"j", "6",
	"c", "7",
	"g", "7",
	"k", "7",
	"q", "7",
	"f", "8",
	"v", "8",
	"b", "9",
	"p", "9",
)
