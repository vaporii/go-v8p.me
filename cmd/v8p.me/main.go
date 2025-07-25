package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vaporii/go-v8p.me/internal/config"
)

func main() {
	conf := config.LoadConfig()

	r := chi.NewRouter()

	log.Println(http.ListenAndServe(conf.ServerAddress, r))
}
