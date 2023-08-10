package main

import (
	log "github.com/shyunku-libraries/go-logger"
	"rotten_world_server/api"
)

func main() {
	log.Info("Starting server...")
	api.Run()
}
