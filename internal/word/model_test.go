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
		err     error
	}{
		{"test T EH1 S T", "test", "T EH1 S T", "101",  nil},
		{"analysis(2) AH0 N AE1 L IH0 S IH0 S", "analysis", "AH0 N AE1 L IH0 S IH0 S", "2500",  nil},
		{"gdp G IY1 D IY1 P IY1 # abbrev", "gdp", "G IY1 D IY1 P IY1", "719",  nil},
		{"tests' T EH1 S T S", "tests'", "T EH1 S T S", "1010",  nil},
		{"a.m. EY2 EH1 M", "a.m.", "EY2 EH1 M", "3",  nil},
		{"computer-generated K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D", "computer-generated", "K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D", "7391462411",  nil},
		{"waah W AA1", "waah", "W AA1", "",  nil},
		{"invalid IH1 N V AH0 L AH0 D ASDF", "invalid", "IH1 N V AH0 L AH0 D ASDF", "2851",  nil},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("Create word from %v", tc.word), func(t *testing.T) {
			t.Parallel()
			wordModel, err := New(tc.input)
			if err != tc.err {
				t.Errorf("got %v, want %v", err, tc.err)
			}
			word, err := wordModel.Word.Value()
			if err != nil {
				t.Error(err)
			}
			if word != word {
				t.Errorf("invalid word. got %s, want %s", word, tc.word)
			}
			if wordModel.Arpabet != tc.arpabet {
				t.Errorf("invalid arpabet. got %s, want %s", wordModel.Arpabet, tc.arpabet)
			}
			number, err := wordModel.Number.Value()
			if err != nil {
				t.Error(err)
			}
			if number != tc.number {
				t.Errorf("invalid number. got %s, want %s", number, tc.number)
			}
		})
	}
}
