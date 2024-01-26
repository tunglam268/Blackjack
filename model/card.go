package model

const (
	PlayerWon = "PLAYER WON"
	DealerWon = "DEALER WON"
	Draw      = "DRAW"
	BlackJack = "BLACKJACK"
)

type Card struct {
	Suit     string
	Value    string
	IsDealer bool
}
