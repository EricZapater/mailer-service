package main

import (
	"log"
	"mailer-service/config"
	"mailer-service/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	server := server.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}