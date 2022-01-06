package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Router(rootFs fs.FS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fileserver := http.FileServer(http.FS(rootFs))

	// Serve index as 404
	r.NotFound(func(res http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/"
		fileserver.ServeHTTP(res, req)
	})

	r.Get("/*", func(res http.ResponseWriter, req *http.Request) {
		requestPath := filepath.Clean(req.URL.Path)
		requestPath = strings.TrimLeft(requestPath, "/")
		if _, err := fs.Stat(rootFs, requestPath); !os.IsNotExist(err) {
			fileserver.ServeHTTP(res, req)
		} else {
			r.NotFoundHandler().ServeHTTP(res, req)
		}
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Get("/number/{query}", ConversionHandler(SearchNumber))
		r.Get("/word/{query}", ConversionHandler(SearchWord))
	})

	return r
}
