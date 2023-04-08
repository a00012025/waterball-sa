package player

import (
	"fmt"
	"strings"

	"github.com/a00012025/waterball-sa/hw-1h/card"
)

type humanPlayer struct {
	basePlayer
}

func (p *humanPlayer) ChooseName() {
	fmt.Printf("Please enter your name: ")
	fmt.Scanln(&p.name)
}

func (p *humanPlayer) TakeTurn(round int) *card.Card {
	if p.exchangeCard != nil && p.exchangeCard.UntilRound == round {
		// needs to exchange back
		fmt.Printf("Exchanging back the cards with %s\n", p.exchangeCard.Opponent.Name())
		p.ExchangeBack()
	}

	if (p.cards == nil) || (len(p.cards) == 0) {
		// no cards to play
		fmt.Printf("No cards to play\n")
		return nil
	}

	fmt.Printf("Your cards:")
	for i, card := range p.cards {
		fmt.Printf(" [%d]%s", i, card.String())
	}
	fmt.Printf("\nPlease choose a card to play: ")

	index := 0
	fmt.Scanln(&index)

	// take out card and return
	card := p.cards[index]
	p.cards = append(p.cards[:index], p.cards[index+1:]...)
	return card
}

func (p *humanPlayer) UseExchangeHands() int {
	fmt.Printf("Please decide whether to exchange hands (y/n): ")
	exchange := ""
	fmt.Scanln(&exchange)
	if strings.ToLower(exchange) == "n" {
		return 0
	}
	fmt.Printf("Please choose a player (1~3 away from you) to exchange hands with: ")
	index := 0
	fmt.Scanln(&index)
	return index
}

// NewHumanPlayer returns a new human player
func NewHumanPlayer() IPlayer {
	return &humanPlayer{}
}
