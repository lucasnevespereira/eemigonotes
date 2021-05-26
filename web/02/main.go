package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func someHandlerFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello"))
}

type Controller struct {
	Message string
}

func (c *Controller) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, c.Message)
}

func main() {
	path, _ := os.Getwd()

	var dir = http.Dir(filepath.Join(path, "public"))

	var fileServer = http.FileServer(dir)
	fmt.Printf("%#v\n", fileServer)

	http.Handle("/", &Controller{"Hello World"})
	http.Handle("/public/", http.StripPrefix("/public", fileServer))
	http.HandleFunc("/hello", someHandlerFunc)
	http.HandleFunc("/fav", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, filepath.Join(path, "public", "favicon.ico"))
	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
