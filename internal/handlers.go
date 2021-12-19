package internal

import (
	"encoding/json"
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func ConversionHandler(db *gorm.DB, searchCol string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		query := chi.URLParam(r, searchCol)

		var words []word.Word
		if err := db.Where(map[string]string{searchCol: query}).Find(&words).Error; err != nil {
			panic(err)
		}

		output, err := json.Marshal(words)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(output)
		if err != nil {
			panic(err)
		}
	}
}

func NumToWordHandler(db *gorm.DB) http.HandlerFunc {
	return ConversionHandler(db, "number")
}

func WordToNumberHandler(db *gorm.DB) http.HandlerFunc {
	return ConversionHandler(db, "word")
}
