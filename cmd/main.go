package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/dynastiateam/backend/repository"
	"github.com/dynastiateam/backend/router"
	"github.com/dynastiateam/backend/services/user"
)

//todo add correct http codes
//todo change os.Getenv to config injection
//todo logging middleware
//todo gracefull shutdown

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_SCHEMA"),
		os.Getenv("DB_SSL")))
	if err != nil {
		log.Fatal(err)
	}

	router := router.New(user.New(repository.New(db)))

	fmt.Println("Started HTTP server on " + os.Getenv("HTTP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), router); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
