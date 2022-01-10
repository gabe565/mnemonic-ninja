package server

import (
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"net/http"
)

type ConversionEntry struct {
	Query string `json:"query"`
	Count int    `json:"count"`
	Guess bool   `json:"guess,omitempty"`

	Words  []*word.WordModel `json:"-"`
	Result []string          `json:"result"`
}

type ConversionResponse struct {
	Query     string             `json:"query"`
	Count     int                `json:"count"`
	QueryType QueryType          `json:"-"`
	Result    []*ConversionEntry `json:"result"`
}

func (response *ConversionResponse) Render(w http.ResponseWriter, r *http.Request) (err error) {
	for _, result := range response.Result {
		for _, w := range result.Words {
			var responseEntry interface{}
			switch response.QueryType {
			case QueryWord:
				responseEntry, err = w.Number.Value()
			case QueryNumber:
				responseEntry, err = w.Word.Value()
			}
			if err != nil {
				return err
			}
			result.Result = append(result.Result, responseEntry.(string))
			if w.Guess {
				result.Guess = w.Guess
			}
		}
		result.Count = len(result.Result)
	}
	response.Count = len(response.Result)
	return err
}
