package card

// Card is a struct that contains a rank and a suit
type Card struct {
	Rank Rank
	Suit Suit
}

// Compare compares two cards and returns -1 if the receiver is less than other. 1 if grater. 0 if equal
func (c *Card) Compare(other *Card) int {
	if c.Rank == other.Rank {
		return compare(int(c.Suit), int(other.Suit))
	}
	return compare(int(c.Rank), int(other.Rank))
}

func (c *Card) String() string {
	return c.Suit.String() + c.Rank.String()
}
