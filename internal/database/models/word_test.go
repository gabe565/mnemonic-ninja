package models

import (
	"fmt"
	"testing"
)

func TestFromCmudict(t *testing.T) {
	testCases := []struct {
		input   string
		word    string
		arpabet string
		number  string
		err     error
	}{
		{"test T EH1 S T", "test", "T EH1 S T", "101", nil},
		{"analysis(2) AH0 N AE1 L IH0 S IH0 S", "analysis", "AH0 N AE1 L IH0 S IH0 S", "2500", nil},
		{"gdp G IY1 D IY1 P IY1 # abbrev", "gdp", "G IY1 D IY1 P IY1", "719", nil},
		{"tests' T EH1 S T S", "tests'", "T EH1 S T S", "1010", nil},
		{"a.m. EY2 EH1 M", "a.m.", "EY2 EH1 M", "3", nil},
		{"computer-generated K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D", "computer-generated", "K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D", "7391462411", nil},
		{"waah W AA1", "waah", "W AA1", "", nil},
		{"invalid IH1 N V AH0 L AH0 D ASDF", "invalid", "IH1 N V AH0 L AH0 D ASDF", "2851", ErrInvalidArpabet},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("Create word from cmudict %v", tc.word), func(t *testing.T) {
			t.Parallel()
			word, err := FromCmudict(tc.input)
			if err != tc.err {
				t.Errorf("unexpected error. got %v, want %v", err, tc.err)
			}
			if word.Word != tc.word {
				t.Errorf("invalid word. got %s, want %s", word.Word, tc.word)
			}
			if word.Arpabet != tc.arpabet {
				t.Errorf("invalid arpabet. got %s, want %s", word.Arpabet, tc.arpabet)
			}
			if word.Number != tc.number {
				t.Errorf("invalid number. got %s, want %s", word.Number, tc.number)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	testCases := []struct {
		input   string
		word    string
		arpabet string
		number  string
		err     error
	}{
		{input: "test", word: "test", number: "101", err: nil},
		{input: "garage", word: "garage", number: "747", err: nil},
		{input: "garages", word: "garages", number: "7470", err: nil},
		{input: "garagez", word: "garagez", number: "7470", err: nil},
		{input: "GARAGES", word: "garages", number: "7470", err: nil},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Create word from user input %s", tc.input), func(t *testing.T) {
			t.Parallel()
			word := FromString(tc.input)
			if word.Word != tc.word {
				t.Errorf("invalid word. got %s, want %s", word.Word, tc.word)
			}
			if word.Arpabet != tc.arpabet {
				t.Errorf("invalid arpabet. got %s, want %s", word.Arpabet, tc.arpabet)
			}
			if word.Number != tc.number {
				t.Errorf("invalid number. got %s, want %s", word.Number, tc.number)
			}
		})
	}
}