package internal

import (
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

type ConversionResponse struct {
	Query  string      `json:"query"`
	Count  int         `json:"count"`
	Result []*word.Word `json:"result"`
}

func (response *ConversionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	response.Query = chi.URLParam(r, "query")
	response.Count = len(response.Result)
	return nil
}

func ConversionHandler(db *gorm.DB, searchCol string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		query := chi.URLParam(r, "query")

		var response ConversionResponse
		if err := db.Where(map[string]string{searchCol: query}).Find(&response.Result).Error; err != nil {
			panic(err)
		}

		err = render.Render(w, r, &response)
		if err != nil {
			panic(err)
		}
	}
}
