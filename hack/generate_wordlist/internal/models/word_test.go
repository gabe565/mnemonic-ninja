package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromCmudict(t *testing.T) {
	testCases := []struct {
		input   []byte
		word    string
		arpabet string
		number  string
		err     error
	}{
		{[]byte("test T EH1 S T"), "test", "T EH1 S T", "101", nil},
		{[]byte("analysis(2) AH0 N AE1 L IH0 S IH0 S"), "analysis", "AH0 N AE1 L IH0 S IH0 S", "2500", nil},
		{[]byte("gdp G IY1 D IY1 P IY1 # abbrev"), "gdp", "G IY1 D IY1 P IY1", "719", nil},
		{[]byte("tests' T EH1 S T S"), "tests'", "T EH1 S T S", "1010", nil},
		{[]byte("a.m. EY2 EH1 M"), "a.m.", "EY2 EH1 M", "3", nil},
		{[]byte("computer-generated K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D"), "computer-generated", "K AH0 M P Y UW1 T ER0 JH EH1 N ER0 EY2 T AH0 D", "7391462411", nil},
		{[]byte("waah W AA1"), "waah", "W AA1", "", nil},
		{[]byte("invalid IH1 N V AH0 L AH0 D ASDF"), "invalid", "IH1 N V AH0 L AH0 D ASDF", "2851", ErrInvalidArpabet},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("Create word from cmudict %v", tc.word), func(t *testing.T) {
			word, err := FromCmudict(tc.input)
			if !assert.ErrorIs(t, err, tc.err) {
				return
			}
			assert.Equal(t, tc.word, word.Word)
			assert.Equal(t, tc.arpabet, word.Arpabet)
			assert.Equal(t, tc.number, word.Number)
		})
	}
}

func Test_truncateAtByte(t *testing.T) {
	type args struct {
		b []byte
		c byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"parenthesis", args{[]byte("testers(2)"), '('}, []byte("testers")},
		{"comment", args{[]byte("G IY1 D IY1 P IY1 # abbrev"), '#'}, []byte("G IY1 D IY1 P IY1 ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncateAtByte(tt.args.b, tt.args.c)
			assert.Equal(t, tt.want, got)
		})
	}
}