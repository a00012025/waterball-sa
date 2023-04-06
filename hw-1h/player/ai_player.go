package player

import (
	"fmt"
	"math/rand"

	"github.com/a00012025/waterball-sa/hw-1h/card"
)

type aiPlayer struct {
	basePlayer
}

func (p *aiPlayer) ChooseName() {
	// random name suffix
	p.name = fmt.Sprintf("AI Player %d", rand.Intn(100))
}

func (p *aiPlayer) TakeTurn(round int) *card.Card {
	if p.exchangeCard != nil && p.exchangeCard.UntilRound == round {
		// needs to exchange back
		p.ExchangeBack()
	}

	if (p.cards == nil) || (len(p.cards) == 0) {
		// no cards to play
		return nil
	}

	index := rand.Intn(len(p.cards))

	// take out card and return
	card := p.cards[index]
	p.cards = append(p.cards[:index], p.cards[index+1:]...)
	return card
}

func (p *aiPlayer) UseExchangeHands() int {
	return rand.Intn(4)
}

// NewAIPlayer returns a new AI player
func NewAIPlayer() IPlayer {
	return &aiPlayer{}
}
