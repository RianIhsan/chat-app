package main

import (
	"github.com/RianIhsan/chat-app/internal/config"
	"github.com/RianIhsan/chat-app/internal/server"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	srv := server.NewServer()

	log.Printf("Server started on port %s\n", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, srv); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
