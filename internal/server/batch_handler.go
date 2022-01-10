package server

import (
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"regexp"
)

var SplitRegex = regexp.MustCompile("[+,; ]+")

func BatchHandler(db *gorm.DB, queryType QueryType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		fullQuery := chi.URLParam(r, "query")
		fullQuery, err = url.QueryUnescape(fullQuery)
		if err != nil {
			panic(err)
		}
		queries := SplitRegex.Split(fullQuery, -1)
		response := ConversionResponse{
			Query: fullQuery,
			QueryType: queryType,
		}
		for _, query := range queries {
			if query == "" {
				continue
			}
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
