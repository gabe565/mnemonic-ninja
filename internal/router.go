package internal

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

const Public = "dist"

func Router(db *gorm.DB, rootFs fs.FS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve index as 404
	r.NotFound(func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, Public+"/index.html")
	})

	// Static Files
	contentFs, err := fs.Sub(rootFs, Public)
	if err != nil {
		panic(err)
	}

	fileserver := http.FileServer(http.FS(contentFs))
	r.Get("/*", func(res http.ResponseWriter, req *http.Request) {
		requestPath := filepath.Join(Public, filepath.Clean("/"+req.URL.Path))
		if _, err := os.Stat(requestPath); !os.IsNotExist(err) {
			fileserver.ServeHTTP(res, req)
		} else {
			r.NotFoundHandler().ServeHTTP(res, req)
		}
	})

	tx := db.Begin(&sql.TxOptions{
		ReadOnly: true,
	})
	r.Get("/api/to/word/{number}", NumToWordHandler(tx))
	r.Get("/api/to/num/{word}", WordToNumberHandler(tx))

	return r
}
