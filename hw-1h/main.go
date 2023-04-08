package main

import (
	"github.com/a00012025/waterball-sa/hw-1h/player"
	"github.com/a00012025/waterball-sa/hw-1h/showdown"
)

func main() {
	// player1 := player.NewHumanPlayer()
	player1 := player.NewAIPlayer()
	// player2 := player.NewHumanPlayer()
	player2 := player.NewAIPlayer()
	player3 := player.NewHumanPlayer()
	player4 := player.NewHumanPlayer()
	game := showdown.NewShowdown([]player.IPlayer{player1, player2, player3, player4})
	game.Play()
}
