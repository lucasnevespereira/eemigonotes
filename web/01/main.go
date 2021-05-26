package main

import (
	"fmt"
	"net/http"
)

func someHandlerFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello"))
}

type Controller struct {
	Message string
}

func (c *Controller) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, c.Message, r.Method)
}

func main() {
	http.HandleFunc("/hello", someHandlerFunc)
	http.Handle("/world", &Controller{"World"})
	http.ListenAndServe(":3000", nil)
}
