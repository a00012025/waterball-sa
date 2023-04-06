package showdown

import (
	"fmt"

	"github.com/a00012025/waterball-sa/hw-1h/card"
	"github.com/a00012025/waterball-sa/hw-1h/deck"
	"github.com/a00012025/waterball-sa/hw-1h/player"
)

// Showdown represents the game
type Showdown struct {
	deck    deck.IDeck
	players []GamePlayer
}

// NewShowdown returns a new Showdown game
func NewShowdown(players []player.IPlayer) Showdown {
	if len(players) != 4 {
		panic("Showdown game requires 4 players")
	}

	gamePlayers := []GamePlayer{}
	for _, p := range players {
		gamePlayers = append(gamePlayers, GamePlayer{
			UseExchangeHands: false,
			Player:           p,
		})
	}
	return Showdown{
		deck:    deck.NewDeck(),
		players: gamePlayers,
	}
}

// Play the game
func (s *Showdown) Play() {
	// Players name themselves
	for i, p := range s.players {
		fmt.Printf("Player %d choosing name...\n", i)
		p.Player.ChooseName()
		fmt.Printf("Player %d finished choosing name\n", i)
	}

	// shuffle cards
	s.deck.Shuffle()

	// all player draw cards
	for i := 0; i < 52; i++ {
		s.players[i%4].Player.DrawCard(s.deck)
	}

	// print all player's cards
	for i, p := range s.players {
		fmt.Printf("Player %d has cards: %v\n", i, p.Player.Cards())
	}

	// play 13 rounds
	for i := 0; i < 13; i++ {
		fmt.Println("=== Round", i+1, "starts ===")

		// all players play and record winner
		winnedPlayer := -1
		var winnedCard *card.Card
		for i, p := range s.players {
			fmt.Printf("Player %d playing...\n", i)

			// check exchange hands ability
			if !p.UseExchangeHands {
				useExchangeHands := p.Player.UseExchangeHands()
				if useExchangeHands > 0 {
					target := (i + useExchangeHands) % 4
					fmt.Printf("Player %d used exchange hands ability with Player %d!\n", i, target)
					p.UseExchangeHands = true
					opponent := s.players[target]
					p.Player.ExchangeHands(i+3, opponent.Player)
				}
			}

			card := p.Player.TakeTurn(i + 1)
			fmt.Printf("Player %d played\n", i)

			if card != nil && (winnedCard == nil || card.Compare(winnedCard) > 0) {
				winnedPlayer = i
				winnedCard = card
			}
		}

		// print winner and add points
		fmt.Println("Player", winnedPlayer, "wins the round!")
		s.players[winnedPlayer].Player.AddPoint()
	}

	// check winner and print
	winner := -1
	for i, p := range s.players {
		if winner == -1 || p.Player.GetPoint() > s.players[winner].Player.GetPoint() {
			winner = i
		}
	}
	fmt.Println("Player", winner, "wins the game!")
}
