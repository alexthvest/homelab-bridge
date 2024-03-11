package main

import (
	"log"

	"github.com/alexthvest/homelab-bridge/pkg/homelab"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	bridge := setupBridge()
	services := setupServices()

	if err := homelab.New(bridge, services).Listen(); err != nil {
		log.Fatalf("failed to listen homelab: %v", err)
	}
}
