package word

import (
	"fmt"
	"testing"
)

func TestNewWord(t *testing.T) {
	testCases := []struct {
		input   string
		word    string
		arpabet string
		number  string
	}{
		{"test T EH1 S T", "test", "T EH1 S T", "101"},
		{"analysis(2) AH0 N AE1 L IH0 S IH0 S", "analysis", "AH0 N AE1 L IH0 S IH0 S", "2500"},
		{"gdp G IY1 D IY1 P IY1 # abbrev", "gdp", "G IY1 D IY1 P IY1", "719"},
		{"tests' T EH1 S T S", "tests'", "T EH1 S T S", "1010"},
		{"a.m. EY2 EH1 M", "a.m.", "EY2 EH1 M", "3"},
		{"computer-generated K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D", "computer-generated", "K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D", "7391462411"},
		{"waah W AA1", "waah", "W AA1", ""},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("Create word from %v", tc.word), func(t *testing.T) {
			t.Parallel()
			word := New(tc.input)
			if *word.Word != tc.word {
				t.Errorf("invalid word. got %s, want %s", *word.Word, tc.word)
			}
			if word.Arpabet != tc.arpabet {
				t.Errorf("invalid arpabet. got %s, want %s", word.Arpabet, tc.arpabet)
			}
			if *word.Number != tc.number {
				t.Errorf("invalid number. got %s, want %s", *word.Number, tc.number)
			}
		})
	}
}
