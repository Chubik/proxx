package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proxx/board"

	"github.com/go-chi/chi/v5"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {

	port := os.Getenv("HOST_PORT")
	if port == "" {
		port = DEFAULT_PORT //default port
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

		// create new handler for flagging operation
		r.Post("/flag/{x}/{y}", func(w http.ResponseWriter, r *http.Request) {
			gs.FlagOperationHandler(w, r, true)
		})

		// create new handler for unflagging operation
		r.Post("/unflag/{x}/{y}", func(w http.ResponseWriter, r *http.Request) {
			gs.FlagOperationHandler(w, r, false)
		})
	})

	r.Get("/", gs.StartGameHandler) // This handler should create a new game, generate a unique ID, and redirect to /game/{player}

	fmt.Printf("Server PROXX is running on http://%s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
