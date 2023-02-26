package main

import (
	"fmt"
	"net/http"
	"strconv"

	f "github.com/fauna/faunadb-go/v4/faunadb"
	mux "github.com/gorilla/mux"
)

func (app *application) api(w http.ResponseWriter, r *http.Request) {
	res, err := app.client.Query(
		f.Map(
			f.Paginate(f.Documents(f.Collection("characters"))),
			f.Lambda("char", f.Get(f.Var("char"))),
		))

	if err != nil {
		panic(err)
	}
	respondWithJSON(w, http.StatusOK, res)
}

func (app *application) getCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	i, err := strconv.Atoi(id)
	res, err := app.client.Query(
		f.Map(
			f.Paginate(f.MatchTerm(f.Index("characters_by_id"), i)),
			f.Lambda("char", f.Get(f.Var("char"))),
		))

	if err != nil {
		panic(err)
	}
	respondWithJSON(w, http.StatusOK, res)
}
