package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	service "github.com/dynastiateam/backend/auth"
	userClient "github.com/dynastiateam/backend/auth/client/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

var (
	// Version is the current version of application
	Version = "0"
	// Branch is the branch this binary was built from
	Branch = "0"
	// Commit is the commit this binary was built from
	Commit = "0"
	// BuildTime is the time this binary was built
	BuildTime = time.Now().Format(time.RFC822)
)

//nolint: funlen
func main() {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("error loading .env file:" + err.Error())
		}
	}
	cfg, err := service.InitConfig()
	if err != nil {
		log.Fatal("failed to init config: " + err.Error())
	}

	log := newLogger(cfg.LogVerbose)

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.SSL))
	if err != nil {
		log.Fatal().Err(err)
	}

	userSrv := userClient.New(cfg.UserServiceHost)
	srv := service.NewService(log, db, userSrv, cfg.JWTSecret)
	handler := service.NewHTTPHandler(srv, log)

	if h, ok := handler.(*chi.Mux); ok {
		h.Get("/user/about", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]string{
				"version": Version,
				"branch":  Branch,
				"commit":  Commit,
				"time":    BuildTime,
			}) //nolint: errcheck
		})
		h.Get("/health", func(w http.ResponseWriter, r *http.Request) {})
	}

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: handler,
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal().Err(err).Msg("error on server shutdown")
		}

		close(signals)
	}()

	log.Info().Msg(fmt.Sprintf("HTTP listener started on :%s @ %s", cfg.HTTPPort, time.Now().Format(time.RFC3339)))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}

func newLogger(verbose bool) (logger *zerolog.Logger) {
	switch verbose {
	case true:
		devLogger := zerolog.New(zerolog.ConsoleWriter{
			NoColor: false,
			Out:     os.Stdout,
		}).Level(zerolog.DebugLevel).With().Timestamp().Logger()
		logger = &devLogger
	default:
		prodLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().Timestamp().Logger()
		logger = &prodLogger
	}
	return logger
}
