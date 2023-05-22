package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proxx/board"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //default port
	}

	gs := board.NewGameServer()

	r := chi.NewRouter()

	host, err := os.Hostname()
	if err != nil {
		host = ""
	}

	r.Route("/game/{player}", func(r chi.Router) {
		r.Get("/", gs.StatusHandler)
		r.Post("/open/{x}/{y}", gs.OpenCellHandler)
		r.Post("/flag/{x}/{y}", gs.FlagCellHandler)
		r.Post("/unflag/{x}/{y}", gs.UnflagCellHandler)
	})

	r.Get("/", gs.StartGameHandler) // This handler should create a new game, generate a unique ID, and redirect to /game/{player}

	fmt.Printf("Server PROXX is running on http://%s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
