package player

// ExchangeCard is a relation between two players
type ExchangeCard struct {
	// UntilRound is the round number until which the exchange is valid
	UntilRound int

	// Opponent is the opponent player
	Opponent IPlayer
}
