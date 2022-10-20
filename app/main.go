package main

import (
	"key-value-data-service/app/config"
	"key-value-data-service/app/server"
	"log"
	"os"
)

func main() {
	config.Load()
	srv := server.NewServer()

	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg != "" {
			config.AppConfig.StorageHost = arg
		}
	}

	log.Printf("server listening on  %s:%d\n", config.AppConfig.Host, config.AppConfig.Port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error: %v, failed to start server on %s:%d", err, config.AppConfig.Host, config.AppConfig.Port)
	}
}
