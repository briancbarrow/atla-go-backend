package main

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/characters", app.searchForCharacter).Queries("search", "{search}")
	r.HandleFunc("/api/characters", app.allCharacters)
	r.HandleFunc("/api/characters/{id}", app.getCharacter)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./docs/public/")))

	// TODO: Add rate limits
	// https://github.com/didip/tollbooth
	// https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g

	// TODO: Add restriction to just GET requests

	// TODO: Add Docs

	//TODO: add route and handler function to be able to search by NAME, then other terms

	// TODO: Add authentication (This might not be needed if we limit to just GET requests, but it would be a good exercise in knowing how to do it)

	// TODO: Add Hosting
	// TODO: Make sure Logging is set up correctly

	return r
}
