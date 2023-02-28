package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Character struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	URL           string   `json:"url"`
	Image         string   `json:"image"`
	Nicknames     []string `json:"nicknames"`
	Aliases       []string `json:"aliases"`
	Nationality   []string `json:"nationality"`
	Ethnicity     []string `json:"ethnicity"`
	FightingStyle []string `json:"fightingStyle"`
	Died          []string `json:"died"`
	HairColor     []string `json:"hair_color"`
	Gender        string   `json:"gender"`
	Weapons       []string `json:"weapons"`
	Profession    []string `json:"profession"`
	Affiliation   []string `json:"affiliation"`
	VoicedBy      []string `json:"voiced_by"`
}

func (app *application) allCharacters(w http.ResponseWriter, r *http.Request) {
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		if err != nil {
			panic(err)
		}

	}
	respondWithJSON(w, http.StatusOK, results)
}

func (app *application) getCharacter(w http.ResponseWriter, r *http.Request) {
	coll := app.mongoClient.Database("Atla-API").Collection("characters")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("id is not an integer")
	}

	var result Character
	err = coll.FindOne(context.TODO(), bson.M{"id": i}).Decode(&result)

	respondWithJSON(w, http.StatusOK, result)
}

func (app *application) getDocs(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/docs", http.StatusSeeOther)
}

func (app *application) searchForCharacter(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	fmt.Println(search)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": search, "$options": "i"}},
			{"nicknames": bson.M{"$regex": search, "$options": "i"}},
			{"aliases": bson.M{"$regex": search, "$options": "i"}},
			{"affiliation": bson.M{"$regex": search, "$options": "i"}},
			{"fightingStyle": bson.M{"$regex": search, "$options": "i"}},
			{"nationality": bson.M{"$regex": search, "$options": "i"}},
			{"profession": bson.M{"$regex": search, "$options": "i"}},
			{"weapons": bson.M{"$regex": search, "$options": "i"}},
			{"ethnicity": bson.M{"$regex": search, "$options": "i"}},
		}})

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		if err != nil {
			panic(err)
		}

	}
	respondWithJSON(w, http.StatusOK, results)
}
