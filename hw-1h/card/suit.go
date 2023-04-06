package card

// Suit is an enum for suits
type Suit int

const (
	// Club .
	Club Suit = iota
	// Diamond .
	Diamond
	// Heart .
	Heart
	// Spade .
	Spade
)

// String method for Suit
func (s Suit) String() string {
	return [...]string{"♣", "♦", "♥", "♠"}[s]
}
