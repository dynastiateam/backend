package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"

	"github.com/dynastiateam/backend"
	"github.com/dynastiateam/backend/repository"
	"github.com/dynastiateam/backend/router"
)

//todo gracefull shutdown
//todo loggin middleware

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("unable to load env file")
	}

	db, err := sql.Open("postgres", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SCHEMA")))
	if err != nil {
		log.Fatal(err)
	}

	router := router.New(backend.New(repository.New(db)))

	go func() {
		if err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), router.Handler); err != nil {
			log.Fatalf("error in ListenAndServe: %s", err)
		}
	}()
	fmt.Println("Started HTTP server on " + os.Getenv("HTTP_PORT"))

	select {}
}
