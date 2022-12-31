package handlers

import (
	"github.com/gabe565/mnemonic-ninja/internal/database/seeds"
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
		request := models.ConversionRequest{QueryType: queryType}
		switch r.Method {
		case http.MethodPost:
			if err := render.Bind(r, &request); err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		case http.MethodGet:
			query := chi.URLParam(r, "query")
			query, err := url.QueryUnescape(query)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			request.Query = query
		}

		response := models.ConversionResponse{
			ConversionRequest: &request,
		}

		<-seeds.Done

		if request.Query != "" {
			for i := len(request.Query); i > 0; i-- {
				entry := models.ConversionEntry{Query: request.Query[0:i]}

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

		if err := render.Render(w, r, &response); err != nil {
			panic(err)
		}
	}
}
