package main

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {

	r := mux.NewRouter()
	r.PathPrefix("/docs").Handler(http.StripPrefix("/docs", http.FileServer(http.Dir("./docs/public/"))))
	r.HandleFunc("/api/characters", app.api)
	r.HandleFunc("/api/characters/{id}", app.getCharacter)
	// catch all route to redirect to docs
	r.PathPrefix("/").Handler(http.HandlerFunc(app.getDocs))
	// TODO: add search route
	// r.HandleFunc("/api/character?search={}", app.searchForCharacter)

	// TODO: Add rate limits

	// TODO: Add restriction to just GET requests

	// TODO: Add authentication (This might not be needed if we limit to just GET requests, but it would be a good exercise in knowing how to do it)

	return r
}
