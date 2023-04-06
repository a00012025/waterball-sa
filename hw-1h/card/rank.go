package card

// Rank is an enum for ranks
type Rank int

const (
	// Two .
	Two Rank = iota
	// Three .
	Three
	// Four .
	Four
	// Five .
	Five
	// Six .
	Six
	// Seven .
	Seven
	// Eight .
	Eight
	// Nine .
	Nine
	// Ten .
	Ten
	// Jack .
	Jack
	// Queen .
	Queen
	// King .
	King
	// Ace .
	Ace
)

// String method for Rank
func (r Rank) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[r]
}
