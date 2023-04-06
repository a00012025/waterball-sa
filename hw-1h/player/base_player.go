package player

import (
	"github.com/a00012025/waterball-sa/hw-1h/card"
	"github.com/a00012025/waterball-sa/hw-1h/deck"
)

type basePlayer struct {
	name         string
	point        int
	cards        []*card.Card
	exchangeCard *ExchangeCard
}

func (p *basePlayer) Name() string {
	return p.name
}

func (p *basePlayer) Point() int {
	return p.point
}

func (p *basePlayer) Cards() []*card.Card {
	return p.cards
}

func (p *basePlayer) SetCards(cards []*card.Card) {
	p.cards = cards
}

func (p *basePlayer) AddPoint() {
	p.point += 1
}

func (p *basePlayer) GetPoint() int {
	return p.point
}

func (p *basePlayer) DrawCard(deck deck.IDeck) {
	card := deck.Draw()
	p.cards = append(p.cards, card)
}

func (p *basePlayer) ExchangeHands(untilRound int, opponent IPlayer) {
	myCards := p.Cards()
	p.SetCards(opponent.Cards())
	opponent.SetCards(myCards)
	p.exchangeCard = &ExchangeCard{
		UntilRound: untilRound,
		Opponent:   opponent,
	}
}

func (p *basePlayer) ExchangeBack() {
	if p.exchangeCard.Opponent != nil {
		myCards := p.Cards()
		opponent := p.exchangeCard.Opponent
		p.SetCards(opponent.Cards())
		opponent.SetCards(myCards)
		p.exchangeCard = nil
	}
}

// IPlayer is an interface for a player
type IPlayer interface {
	Name() string
	Point() int
	Cards() []*card.Card
	SetCards(cards []*card.Card)
	AddPoint()
	GetPoint() int
	DrawCard(deck deck.IDeck)
	ExchangeHands(untilRound int, opponent IPlayer)
	ExchangeBack()

	ChooseName()
	TakeTurn(round int) *card.Card
	UseExchangeHands() int
}
