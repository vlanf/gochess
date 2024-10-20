package main

import (
	"fmt"

	"github.com/vlanf/gochess/internal/game"
)

func main() {
	game := game.New()
	fmt.Println(game.Board)
}
