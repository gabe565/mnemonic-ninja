package models

import (
	"fmt"
	"testing"
)

func TestWhereColumn(t *testing.T) {
	testCases := []struct {
		queryType QueryType
		expect    string
	}{
		{Word, "word"},
		{Number, "number"},
		{99, "QueryType(99)"},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("Get where column for QueryType %v", tc.queryType), func(t *testing.T) {
			t.Parallel()
			got := tc.queryType.WhereColumn()
			if got != tc.expect {
				t.Errorf("invalid where column. got %s, want %s", got, tc.expect)
			}
		})
	}
}

func TestDistinctColumn(t *testing.T) {
	testCases := []struct {
		queryType QueryType
		expect    string
	}{
		{Word, "word"},
		{Number, "number"},
		{99, "QueryType(99)"},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("Get where column for QueryType %v", tc.queryType), func(t *testing.T) {
			t.Parallel()
			got := tc.queryType.WhereColumn()
			if got != tc.expect {
				t.Errorf("invalid where column. got %s, want %s", got, tc.expect)
			}
		})
	}
}
