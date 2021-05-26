package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// RootHandler handles path /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusResetContent)
	fmt.Fprintf(w, "hello, world: %s %s\n", r.Method, r.URL)
}

func main() {
	const assetsPath = "public"

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var (
		public        = filepath.Join(cwd, assetsPath)
		publicDir     = http.Dir(public)
		faviconPath   = filepath.Join(public, "favicon.ico")
		publicHandler = http.FileServer(publicDir)
		publicPrefix  = "/" + assetsPath
	)

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, faviconPath)
	})
	http.Handle(publicPrefix+"/", http.StripPrefix(publicPrefix, publicHandler))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
