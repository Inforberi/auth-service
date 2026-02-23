package main

import (
	"log"

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

	// a := app.NewApp(logg)

}
