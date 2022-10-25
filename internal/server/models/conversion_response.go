package models

import (
	"net/http"
)

type ConversionRequest struct {
	Query     string    `json:"query"`
	QueryType QueryType `json:"-"`
}

func (c *ConversionRequest) Bind(r *http.Request) error {
	return nil
}

type ConversionResponse struct {
	*ConversionRequest
	Count  int                `json:"count"`
	Result []*ConversionEntry `json:"result"`
}

func (response *ConversionResponse) Render(w http.ResponseWriter, r *http.Request) (err error) {
	for _, result := range response.Result {
		for _, w := range result.Words {
			switch response.QueryType {
			case Word:
				result.Result = append(result.Result, w.Number)
			case Number:
				result.Result = append(result.Result, w.Word)
			}
			if w.Guess {
				result.Guess = w.Guess
			}
		}
	}
	response.Count = len(response.Result)
	return err
}
