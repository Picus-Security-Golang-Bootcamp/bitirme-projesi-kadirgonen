package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"HW/app/migration"
	"HW/app/router"
	"HW/config"
	db "HW/pkg/database_handler"
	"HW/pkg/httpserver"
	"HW/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	// Configuration
	cfg, err := config.LoadConfig("./config/config-local")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	//DB
	DB := db.Connect(cfg)
	DB.AutoMigrate()
	//logger
	l := logger.New(cfg.Logger.Level)
	//Migration
	if cfg.Migration.AutoMigrate {
		err := migration.Execute(db.Connect(cfg))
		if err != nil {
			l.Fatal(fmt.Errorf("app - Run - migration.Execute: %w", err))
		}
	}
	// GIN & router
	gin.SetMode(cfg.Gin.GINMode)
	handler := gin.New()
	router.NewRouter(handler, l, cfg, db.Connect(cfg))

	// HTTP server
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Http.HTTPPort))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
