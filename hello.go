package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	fauna "github.com/fauna/faunadb-go/v4/faunadb"
	env "github.com/joho/godotenv"
)

type User struct {
	Name string `fauna:"name"`
}

func main() {

	fmt.Println("Hello, world.")

	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)
}

func servePage(writer http.ResponseWriter, reqest *http.Request) {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	endpoint := os.Getenv("ENDPOINT")
	secretKey := os.Getenv("SECRET")
	client := fauna.NewFaunaClient(
		secretKey,
		fauna.Endpoint(endpoint),
		// NOTE: use the correct endpoint for your database's Region Group.
	)
	// fmt.Println(client)
	// 355977710100545620
	res, err := client.Query(fauna.Get(fauna.Ref(fauna.Collection("Characters"), "355977710100545620")))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", res)
	// var user User

	// if err := res.At(fauna.ObjKey("data")).Get(&user); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)

}
