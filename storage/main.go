package main

import (
	"key-value-data-service/storage/config"
	"key-value-data-service/storage/server"
	"log"
)

func main() {
	config.Load()
	srv := server.NewServer()
	log.Printf("server listening on  %s:%d\n", config.AppConfig.Host, config.AppConfig.Port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error: %v, failed to start server on %s:%d", err, config.AppConfig.Host, config.AppConfig.Port)
	}
}
