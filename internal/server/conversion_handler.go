package server

import (
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"regexp"
)

type QueryType int

const (
	QueryWord QueryType = iota
	QueryNumber
)

func (col QueryType) WhereColumn() string {
	switch col {
	case QueryWord:
		return "word"
	case QueryNumber:
		return "number"
	}
	return ""
}

func (col QueryType) DistinctColumn() string {
	switch col {
	case QueryWord:
		return "number"
	case QueryNumber:
		return "word"
	}
	return ""
}

var SplitRegex = regexp.MustCompile("[+,; ]+")

type ConversionEntry struct {
	Query string `json:"query"`
	Count int    `json:"count"`
	Guess bool   `json:"guess,omitempty"`

	Words  []*word.WordModel `json:"-"`
	Result []string          `json:"result"`
}

type ConversionResponse struct {
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

func ConversionHandler(db *gorm.DB, queryType QueryType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		fullQuery := chi.URLParam(r, "query")
		queries := SplitRegex.Split(fullQuery, -1)
		response := ConversionResponse{QueryType: queryType}
		for _, query := range queries {
			entry := ConversionEntry{Query: query}

			err = db.Distinct(queryType.DistinctColumn()).
				Where(map[string]interface{}{queryType.WhereColumn(): entry.Query}).
				Find(&entry.Words).Error
			if err != nil {
				panic(err)
			}

			if queryType == QueryWord && len(entry.Words) == 0 {
				w, err := word.FromString(query)
				if err != nil {
					panic(err)
				}
				if w.Word.Valid {
					entry.Words = append(entry.Words, w)
				}
			}
			response.Result = append(response.Result, &entry)
		}

		err = render.Render(w, r, &response)
		if err != nil {
			panic(err)
		}
	}
}
