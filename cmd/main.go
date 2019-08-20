package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/valyala/fasthttp"

	"github.com/dynastiateam/backend/repository"
	"github.com/dynastiateam/backend/router"
	"github.com/dynastiateam/backend/services/user"
)

//todo logging middleware
//todo gracefull shutdown

func main() {
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

	go func() {
		if err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), router.Handler); err != nil {
			log.Fatalf("error in ListenAndServe: %s", err)
		}
	}()
	fmt.Println("Started HTTP server on " + os.Getenv("HTTP_PORT"))

	select {}
}
