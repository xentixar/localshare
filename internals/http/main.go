// Package http handles http related stuffs
package http

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

func Start(port int, password string, files []string) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if password != "" {
			if addPasswordProtection(password, w, r) {
				_, _ = w.Write([]byte(createRootPage(files)))
				return
			} else {
				return
			}
		}

		_, _ = w.Write([]byte(createRootPage(files)))
	})

	mux.HandleFunc("GET /download", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Query().Get("path")
		if path == "" {
			http.Error(w, "Missing 'path' parameter", http.StatusBadRequest)
			return
		}

		if password != "" {
			if addPasswordProtection(password, w, r) {
				os.Exit(1)
				download(path, files, w, r)
				return
			} else {
				return
			}
		}

		download(path, files, w, r)
	})

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Server failed: %v", err)
	}
}

func addPasswordProtection(password string, w http.ResponseWriter, r *http.Request) bool {
	user, pass, ok := r.BasicAuth()

	if !ok || user != "localshare" || pass != password {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	return true
}

func createRootPage(files []string) string {
	html := ""
	for _, file := range files {
		html = html + fmt.Sprintf("<a href='/download?path=%v'>%v</a>", file, filepath.Base(file)) + "<br/>"
	}
	return html
}

func download(path string, files []string, w http.ResponseWriter, r *http.Request) {
	if !slices.Contains(files, path) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	extension := filepath.Ext(path)
	contentType := mime.TypeByExtension(extension)

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(path))

	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	http.ServeFile(w, r, path)
}
