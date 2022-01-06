package internal

import (
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type SearchCol int

const (
	SearchWord SearchCol = iota
	SearchNumber
)

func (col SearchCol) Column() string {
	switch col {
	case SearchWord:
		return "word"
	case SearchNumber:
		return "number"
	}
	return ""
}

type ConversionResponse struct {
	Query  string        `json:"query"`
	Count  int           `json:"count"`
	Result []interface{} `json:"result"`
}

func (response ConversionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	response.Query = chi.URLParam(r, "query")
	response.Count = len(response.Result)
	return nil
}

func ConversionHandler(searchCol SearchCol) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		query := chi.URLParam(r, "query")

		var found []*word.WordModel
		err = Db.Where(map[string]interface{}{searchCol.Column(): query}).Find(&found).Error
		if err != nil {
			panic(err)
		}

		var response ConversionResponse
		for _, e := range found {
			var responseEntry interface{}
			switch searchCol {
			case SearchWord:
				responseEntry, err = e.ToWordResponse()
			case SearchNumber:
				responseEntry, err = e.ToNumResponse()
			}
			if err != nil {
				panic(err)
			}
			response.Result = append(response.Result, responseEntry)
		}
		err = render.Render(w, r, response)
		if err != nil {
			panic(err)
		}
	}
}
