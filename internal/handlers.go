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
	Query     string            `json:"query"`
	Count     int               `json:"count"`

	Words     []*word.WordModel `json:"-"`
	SearchCol SearchCol         `json:"-"`

	Result    []interface{}     `json:"result"`
}

func (response *ConversionResponse) Render(w http.ResponseWriter, r *http.Request) (err error) {
	for _, e := range response.Words {
		var responseEntry interface{}
		switch response.SearchCol {
		case SearchWord:
			responseEntry, err = e.Number.Value()
		case SearchNumber:
			responseEntry, err = e.Word.Value()
		}
		if err != nil {
			return err
		}
		response.Result = append(response.Result, responseEntry)
	}
	response.Count = len(response.Result)
	return err
}

func ConversionHandler(searchCol SearchCol) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		response := &ConversionResponse{
			Query: chi.URLParam(r, "query"),
			SearchCol: searchCol,
		}

		err = Db.Where(map[string]interface{}{searchCol.Column(): response.Query}).Find(&response.Words).Error
		if err != nil {
			panic(err)
		}

		err = render.Render(w, r, response)
		if err != nil {
			panic(err)
		}
	}
}
