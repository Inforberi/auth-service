package main

import (
	"log"
	"net/http"

	"github.com/inforberi/auth-service/internal/app"
	"github.com/inforberi/auth-service/internal/config"
	"github.com/inforberi/auth-service/internal/logger"
)

func main() {
	// ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Println("Config init success")

	logg := logger.NewLogger(cfg)
	logg.Info("init logger success")

	application, err := app.NewApp(cfg, logg)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: application.Router,
	}

	log.Fatal(server.ListenAndServe())

}
