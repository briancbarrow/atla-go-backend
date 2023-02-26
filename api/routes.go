package main

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", app.getDocs)
	r.PathPrefix("/docs").Handler(http.StripPrefix("/docs", http.FileServer(http.Dir("./docs/public/"))))
	r.HandleFunc("/api", app.api)
	r.HandleFunc("/api/{id}", app.getCharacter)

	return r
}
