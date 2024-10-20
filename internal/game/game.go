package game

import (
	"github.com/vlanf/gochess/internal/board"
)

type Game struct {
	Board *board.Board
	Timer int // TODO
}

func New() Game {
	var game Game
	game.Board = board.New()
	return game
}
