package server

import (
	"github.com/RianIhsan/chat-app/internal/chat"
	"github.com/gorilla/mux"
)

func NewServer() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ws", chat.HandleConnections)

	go chat.Hub2.Run()

	return r
}
