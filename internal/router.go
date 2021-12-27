package internal

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

func Router(db *gorm.DB, rootFs fs.FS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve index as 404
	r.NotFound(func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "/index.html")
	})

	fileserver := http.FileServer(http.FS(rootFs))
	r.Get("/*", func(res http.ResponseWriter, req *http.Request) {
		requestPath := filepath.Clean("/"+req.URL.Path)
		if _, err := fs.Stat(rootFs, requestPath); !os.IsNotExist(err) {
			fileserver.ServeHTTP(res, req)
		} else {
			r.NotFoundHandler().ServeHTTP(res, req)
		}
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(render.SetContentType(render.ContentTypeJSON))

		tx := db.Begin(&sql.TxOptions{
			ReadOnly: true,
		})

		r.Get("/number/{query}", ConversionHandler(tx, "number", "word"))
		r.Get("/word/{query}", ConversionHandler(tx, "word", "number"))
	})

	return r
}
