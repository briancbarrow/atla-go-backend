package main

import (
	mux "github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api", app.api)
	r.HandleFunc("/api/{id}", app.getCharacter)

	return r
}
