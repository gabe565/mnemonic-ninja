package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
			got := tc.queryType.WhereColumn()
			assert.Equal(t, tc.expect, got)
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
			got := tc.queryType.WhereColumn()
			assert.Equal(t, tc.expect, got)
		})
	}
}
