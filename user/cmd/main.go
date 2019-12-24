package main

import (
	"log"
	"os"
	"time"

	//datadog "github.com/EdisonJunior/stark-common/datadog/v1"
	//logger "github.com/EdisonJunior/stark-common/logger/v1"
	//
	//service "github.com/EdisonJunior/stark-reference-service"
	//
	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/go-chi/chi"
	"github.com/joho/godotenv"
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

	//cfg, err := service.InitConfig(os.Getenv("APP_ENV"), awsSession)
	//if err != nil {
	//	log.Fatal("failed to init config: " + err.Error())
	//}
	//
	//log := logger.New(cfg.LogVerbose)
	//
	//var ddog *datadog.Client = nil
	//if cfg.DataDog.Enabled {
	//	var err error
	//	ddog, err = datadog.New(cfg.DataDog.Host, cfg.DataDog.Port, cfg.DataDog.Namespace)
	//	if err != nil {
	//		log.Error().Err(err).Msg("failed to create datadog client")
	//		log.Warn().Msg("datadog client is not enabled, metrics will not be collected")
	//	}
	//}
	//
	//srv := service.NewService()
	//srv = service.NewLoggingMiddleware(log, srv)
	//srv = service.NewInstrumentingMiddleware(ddog, srv)
	//handler := service.NewHTTPHandler(srv, log)
	//
	//if h, ok := handler.(*chi.Mux); ok {
	//	h.Get("/about", func(w http.ResponseWriter, r *http.Request) {
	//		json.NewEncoder(w).Encode(map[string]string{
	//			"version": Version,
	//			"branch":  Branch,
	//			"commit":  Commit,
	//			"time":    BuildTime,
	//		}) //nolint: errcheck
	//	})
	//	h.Get("/health", func(w http.ResponseWriter, r *http.Request) {})
	//}
	//
	//server := &http.Server{
	//	Addr:    ":" + cfg.HTTPPort,
	//	Handler: handler,
	//}
	//
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	//go func() {
	//	<-signals
	//	if err := server.Shutdown(context.Background()); err != nil {
	//		log.Fatal().Err(err).Msg("error on server shutdown")
	//	}
	//
	//	close(signals)
	//}()
	//
	//log.Info().Msg(fmt.Sprintf("HTTP listener started on :%s @ %s", cfg.HTTPPort, time.Now().Format(time.RFC3339)))
	//if err := server.ListenAndServe(); err != nil {
	//	log.Fatal().Err(err)
	//}
}
