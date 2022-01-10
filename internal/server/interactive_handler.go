package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"net/url"
)

func InteractiveHandler(db *gorm.DB) http.HandlerFunc {
	queryType := QueryNumber
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		query := chi.URLParam(r, "query")
		query, err = url.QueryUnescape(query)
		if err != nil {
			panic(err)
		}
		response := ConversionResponse{
			Query: query,
			QueryType: queryType,
		}

		if query != "" {
			for i := len(query); i > 0; i-- {
				entry := ConversionEntry{Query: query[0:i]}

				err = db.Distinct(queryType.DistinctColumn()).
					Where(map[string]interface{}{queryType.WhereColumn(): entry.Query}).
					Find(&entry.Words).Error
				if err != nil {
					panic(err)
				}

				if len(entry.Words) > 0 {
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
