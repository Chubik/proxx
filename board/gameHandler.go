package board

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Default values for game parameters
const (
	defaultWidth  = 10
	defaultHeight = 10
	defaultMines  = 10

	// errors text constants
	ERROR_WRONG_COORDS_FORMAT = "wrong coords format"
	ERROR_GAME_NOT_FOUND      = "game not found"
)

func (gs *GameServer) StatusHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join("assets", "status.html"))
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, game); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (gs *GameServer) OpenCellHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")
	xStr := chi.URLParam(r, "x")
	yStr := chi.URLParam(r, "y")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	if err := game.OpenCell(x, y); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}

func (gs *GameServer) FlagCellHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")
	xStr := chi.URLParam(r, "x")
	yStr := chi.URLParam(r, "y")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	if err := game.FlagCell(x, y); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}

func (gs *GameServer) UnflagCellHandler(w http.ResponseWriter, r *http.Request) {
	player := chi.URLParam(r, "player")
	xStr := chi.URLParam(r, "x")
	yStr := chi.URLParam(r, "y")

	gs.mu.Lock()
	game, ok := gs.games[player]
	gs.mu.Unlock()
	if !ok {
		log.Println(ERROR_GAME_NOT_FOUND)
		http.NotFound(w, r)
		return
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, ERROR_WRONG_COORDS_FORMAT, http.StatusBadRequest)
		return
	}

	if err := game.UnflagCell(x, y); err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}

func (gs *GameServer) StartGameHandler(w http.ResponseWriter, r *http.Request) {
	// generate unique player id
	player := uuid.NewString()

	// game creation
	game, err := NewGame(defaultWidth, defaultHeight, defaultMines, player)
	if err != nil {
		log.Printf("error game ID %s: %v \n", player, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add game to the server storage
	gs.mu.Lock()
	gs.games[player] = game
	gs.mu.Unlock()

	// Redirect to the game status page
	http.Redirect(w, r, fmt.Sprintf("/game/%s", player), http.StatusFound)
}
