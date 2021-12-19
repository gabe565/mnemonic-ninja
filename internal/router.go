package internal

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tx := db.Begin(&sql.TxOptions{
		ReadOnly: true,
	})
	r.Get("/api/to/word/{number}", NumToWordHandler(tx))
	r.Get("/api/to/num/{word}", WordToNumberHandler(tx))

	return r
}
