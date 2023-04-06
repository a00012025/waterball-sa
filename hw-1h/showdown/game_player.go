package showdown

import "github.com/a00012025/waterball-sa/hw-1h/player"

// GamePlayer is a struct that contains a player and a flag to indicate whether the player has exchanged hands
type GamePlayer struct {
	UseExchangeHands bool
	Player           player.IPlayer
}
