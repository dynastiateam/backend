package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/dynastiateam/backend/config"
	"github.com/dynastiateam/backend/logging"
	"github.com/dynastiateam/backend/router"
	"github.com/dynastiateam/backend/services/auth"
	"github.com/dynastiateam/backend/services/user"
)

//todo gracefull shutdown

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	log := logging.NewLogger(cfg.Debug)

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.SSL))
	if err != nil {
		log.Fatal().Err(err)
	}

	userService := user.NewService(log, db)
	userHandler := user.NewHandler(log, userService)

	authService := auth.NewService(log, db, userService, cfg.JWTSecret)
	authHandler := auth.NewHandler(log, authService)

	r := router.New(cfg.JWTSecret, cfg.HTTPPort, log, userHandler.Routes(), authHandler.Routes())

	log.Info().Msg(fmt.Sprintf("Started HTTP server on %s @ %s", cfg.HTTPPort, time.Now().Format("02/01/2006 15:04:05")))
	if err := r.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}
