# Proxx Game

## Game Description
Proxx is a game where the player is supposed to reveal cells on an NxN board, trying to avoid black holes. Each cell not containing a black hole displays a count of adjacent black holes. Diagonal cells are also considered as adjacent. The game automatically reveals surrounding cells when a cell with zero adjacent black holes is clicked.

This repository contains the game logic and user interface, written in Go. The game can be played through a web browser.

## Game Features
- NxN board representation
- Placement of black holes in random locations
- Counting the number of adjacent black holes for each cell
- Automatic revealing of surrounding cells with zero adjacent black holes
- User interface implemented with golang HTML templates
- In default we have 10x10 board and 10 holes

## Getting Started
To run the game locally, follow these steps:

1. Clone this repository to your local machine.
2. Make sure you have Go installed on your system.
3. Navigate to the cloned repository directory.
4. Run the following command to start the game server:
   ```
   go run .
   ```
5. Open a web browser and visit `http://localhost:8080` to play the game.

Доданий новий пункт в опис README для запуску проекту з використанням Docker Compose:

### Running with Docker Compose

To run the project using Docker Compose, follow these steps:

1. Make sure you have Docker installed on your machine.
2. Create a file named `.env` in the root directory of the project.
3. Open the `.env` file and set the desired values for the ports. For example:
   ```plaintext
   HOST_PORT=8080
   CONTAINER_PORT=8080
   ```
4. Run the following command to start the containers:
   ```plaintext
   docker-compose -f ./docker-compose.yml --env-file .env up --build
   ```
5. Wait for the containers to build and start. Once everything is ready, you can access the Proxx game UI by opening a web browser and navigating to `http://localhost:8080`.

Make sure to adjust the `HOST_PORT` and `CONTAINER_PORT` values in the `.env` file to the desired port numbers. These ports will be used to access the Proxx game UI in the browser.

Please let me know if there's anything else I can help you with!

## Running Tests
To run all tests for the game, use the following command:
```
go test ./... -v
```

## Compiling to Binary
To compile the game into a binary executable, use the following commands:

- For Linux:
  ```
  GOOS=linux GOARCH=amd64 go build -o proxx-linux
  ```
  The compiled binary will be named `proxx-linux`.

- For Windows:
  ```
  GOOS=windows GOARCH=amd64 go build -o proxx-windows.exe
  ```
  The compiled binary will be named `proxx-windows.exe`.

## User Interface
The user interface for the Proxx game is implemented using golang HTML templates. The main and only one template file is located in the `assets/status.html` file. The template includes dynamic placeholders that are filled with game data when rendered.

To access the game status page, open a web browser and visit `http://localhost:8080/game/{playerID}`, where `{playerID}` is a unique identifier for the player.

The game status page displays the game board, showing the covered and revealed cells. The player can click on covered cells to reveal their content. The game state is also displayed, indicating whether the game is in progress, won, or lost.
