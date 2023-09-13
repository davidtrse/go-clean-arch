package main

import (
	"flag"
	"log"
	"os"

	"github.com/davidtrse/go-clean-arch/config"
	"github.com/davidtrse/go-clean-arch/internal/server"
)

func main() {
	// Load configuration
	configPath := flag.String("config", "./config", "config folder path")
	flag.Parse()

	log.Printf("ENV: %s", os.Getenv("APPENV"))
	log.Printf("Config path: %s", *configPath)

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	config.AppConfig = cfg

	// Run server
	server.Run()
}
