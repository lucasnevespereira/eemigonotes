package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func helloHtml(w http.ResponseWriter, r *http.Request) {
}



func handlePost(w http.ResponseWriter, r *http.Request) {
}



func handlePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Printf("page incorrecte: %s \n", page)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Je suis dans la page %d \n", pageInt)
}



func main() {
	// Allouer un routeur Gorilla.
	router := mux.NewRouter()
	// Routes filtrées par les méthodes.
	const MethodEEMI = "EEMI"
	router.HandleFunc("/hello", helloHtml).
		Methods(http.MethodGet, MethodEEMI)
	router.HandleFunc("/", handlePost).
		Methods(http.MethodPost)
	// Route avec paramètres et validation.
	router.Path("/page/{page:[0-9]+}").
		HandlerFunc(handlePage).
		Methods(http.MethodGet)
	http.Handle("/", router)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
