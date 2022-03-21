package handlers

import (
	"github.com/gabe565/mnemonic-ninja/internal/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"net/url"
)

func InteractiveHandler(db *gorm.DB) http.HandlerFunc {
	queryType := models.Number
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		query := chi.URLParam(r, "query")
		query, err = url.QueryUnescape(query)
		if err != nil {
			panic(err)
		}
		response := models.ConversionResponse{
			Query:     query,
			QueryType: queryType,
		}

		if query != "" {
			for i := len(query); i > 0; i-- {
				entry := models.ConversionEntry{Query: query[0:i]}

				result := db.Distinct(queryType.DistinctColumn()).
					Where(map[string]any{queryType.WhereColumn(): entry.Query}).
					Find(&entry.Words)
				if result.Error != nil {
					panic(result.Error)
				}

				entry.Count = result.RowsAffected

				if entry.Count > 0 {
					response.Result = append(response.Result, &entry)
				}
			}
		}

		err = render.Render(w, r, &response)
		if err != nil {
			panic(err)
		}
	}
}
