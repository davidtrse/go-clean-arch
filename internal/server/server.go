package server

import (
	"log"
	"net/http"

	"github.com/davidtrse/go-clean-arch/config"
	"github.com/davidtrse/go-clean-arch/pkg/connection"
	"github.com/gin-gonic/gin"
)

func Run() {
	// create connection with database
	connection, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	engine.RedirectTrailingSlash = false
	api := engine.Group("/api")

	NewRouter(api, connection)

	// Listen HTTP Server
	srv := &http.Server{
		Addr:    config.AppConfig.Server.Port,
		Handler: engine,
	}

	// Start HTTP Server
	log.Printf("HTTP server listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}
