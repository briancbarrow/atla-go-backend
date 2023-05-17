package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	cursor, err := coll.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"id", 1}}))
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

	// use the FindOne query from mongodb
}

func (app *application) searchCharacterByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println(name)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": name, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByNicknames(w http.ResponseWriter, r *http.Request) {
	nicknames := r.URL.Query().Get("nicknames")
	fmt.Println(nicknames)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"nicknames": bson.M{"$regex": nicknames, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByAliases(w http.ResponseWriter, r *http.Request) {
	aliases := r.URL.Query().Get("aliases")
	fmt.Println(aliases)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"aliases": bson.M{"$regex": aliases, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByAffiliation(w http.ResponseWriter, r *http.Request) {
	affiliation := r.URL.Query().Get("affiliation")
	fmt.Println(affiliation)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"affiliation": bson.M{"$regex": affiliation, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByFightingStyle(w http.ResponseWriter, r *http.Request) {
	fightingStyle := r.URL.Query().Get("fightingStyle")
	fmt.Println(fightingStyle)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"fightingStyle": bson.M{"$regex": fightingStyle, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByNationality(w http.ResponseWriter, r *http.Request) {
	nationality := r.URL.Query().Get("nationality")
	fmt.Println(nationality)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"nationality": bson.M{"$regex": nationality, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByProfession(w http.ResponseWriter, r *http.Request) {
	profession := r.URL.Query().Get("profession")
	fmt.Println(profession)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"profession": bson.M{"$regex": profession, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByWeapons(w http.ResponseWriter, r *http.Request) {
	weapons := r.URL.Query().Get("weapons")
	fmt.Println(weapons)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"weapons": bson.M{"$regex": weapons, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

func (app *application) searchCharacterByEthnicity(w http.ResponseWriter, r *http.Request) {
	ethnicity := r.URL.Query().Get("ethnicity")
	fmt.Println(ethnicity)
	coll := app.mongoClient.Database("Atla-API").Collection("characters")
	cursor, err := coll.Find(context.TODO(), bson.M{
		"$or": []bson.M{
			{"ethnicity": bson.M{"$regex": ethnicity, "$options": "i"}},
		}},
		options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusOK, results)

}

// func (app *application) getDocs(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/docs", http.StatusSeeOther)
// }

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
		}}, options.Find().SetSort(bson.D{{"id", 1}}))

	var results []Character
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// for _, result := range results {
	// 	cursor.Decode(&result)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }
	respondWithJSON(w, http.StatusOK, results)
}
