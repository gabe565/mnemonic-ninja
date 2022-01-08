package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"
)

func Router(db *gorm.DB, rootFs fs.FS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	fileserver := http.FileServer(http.FS(rootFs))

	// Serve index as 404
	r.NotFound(func(res http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/"
		fileserver.ServeHTTP(res, req)
	})

	r.Get("/*", func(res http.ResponseWriter, req *http.Request) {
		requestPath := strings.TrimLeft(req.URL.Path, "/")
		if _, err := fs.Stat(rootFs, requestPath); !os.IsNotExist(err) {
			fileserver.ServeHTTP(res, req)
		} else {
			r.NotFoundHandler().ServeHTTP(res, req)
		}
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(httprate.LimitByIP(60, time.Minute))
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Get("/number/{query}", ConversionHandler(db, QueryNumber))
		r.Get("/word/{query}", ConversionHandler(db, QueryWord))
	})

	return r
}
