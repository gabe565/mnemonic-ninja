package handlers

import (
	models2 "github.com/gabe565/mnemonic-ninja/internal/database/models"
	"github.com/gabe565/mnemonic-ninja/internal/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"regexp"
)

var SplitRegex = regexp.MustCompile("[+,; ]+")

func BatchHandler(db *gorm.DB, queryType models.QueryType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		fullQuery := chi.URLParam(r, "query")
		fullQuery, err = url.QueryUnescape(fullQuery)
		if err != nil {
			panic(err)
		}
		queries := SplitRegex.Split(fullQuery, -1)
		response := models.ConversionResponse{
			Query:     fullQuery,
			QueryType: queryType,
		}
		for _, query := range queries {
			if query == "" {
				continue
			}
			entry := models.ConversionEntry{Query: query}

			result := db.Distinct(queryType.DistinctColumn()).
				Where(map[string]any{queryType.WhereColumn(): entry.Query}).
				Find(&entry.Words)
			if result.Error != nil {
				panic(result.Error)
			}

			entry.Count = result.RowsAffected

			if queryType == models.Word && len(entry.Words) == 0 {
				w := models2.FromString(query)
				entry.Words = append(entry.Words, w)
			}
			response.Result = append(response.Result, &entry)
		}

		err = render.Render(w, r, &response)
		if err != nil {
			panic(err)
		}
	}
}
