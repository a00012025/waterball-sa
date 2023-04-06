package deck

import (
	"math/rand"

	"github.com/a00012025/waterball-sa/hw-1h/card"
)

// IDeck is an interface for a deck of cards
type IDeck interface {
	Shuffle()
	Draw() *card.Card
}

// NewDeck returns a new deck of cards
func NewDeck() IDeck {
	return &deck{}
}

// deck is a struct that contains a slice of cards
type deck struct {
	cards []*card.Card
}

// Shuffle initializes all cards and shuffles the deck
func (d *deck) Shuffle() {
	// init all cards
	d.cards = make([]*card.Card, 52)
	for i := 0; i < 52; i++ {
		d.cards[i] = &card.Card{
			Rank: card.Rank(i % 13),
			Suit: card.Suit(i / 13),
		}
	}

	// shuffle
	for i := 0; i < 52; i++ {
		r := rand.Intn(52)
		d.cards[i], d.cards[r] = d.cards[r], d.cards[i]
	}
}

func (d *deck) Draw() *card.Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}
