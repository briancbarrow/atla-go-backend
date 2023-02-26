package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	f "github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/joho/godotenv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	client   *f.FaunaClient
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := f.NewFaunaClient(os.Getenv("FAUNA_SECRET"))
	fmt.Printf("client: %v", client)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		client:   client,
	}
	srv := &http.Server{
		ErrorLog: errorLog,
		Handler:  app.routes(),
		Addr:     ":4444",
	}

	infoLog.Printf("Starting server on %s", ":4444")
	err = srv.ListenAndServe()
	infoLog.Printf("Server stopped")
	errorLog.Fatal(err)
}
