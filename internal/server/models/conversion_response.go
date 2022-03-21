package models

import (
	"net/http"
)

type ConversionResponse struct {
	Query     string             `json:"query"`
	Count     int                `json:"count"`
	QueryType QueryType          `json:"-"`
	Result    []*ConversionEntry `json:"result"`
}

func (response *ConversionResponse) Render(w http.ResponseWriter, r *http.Request) (err error) {
	for _, result := range response.Result {
		for _, w := range result.Words {
			var responseEntry any
			switch response.QueryType {
			case Word:
				responseEntry = w.Number
			case Number:
				responseEntry = w.Word
			}
			if err != nil {
				return err
			}
			result.Result = append(result.Result, responseEntry.(string))
			if w.Guess {
				result.Guess = w.Guess
			}
		}
	}
	response.Count = len(response.Result)
	return err
}
