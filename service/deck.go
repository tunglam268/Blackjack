package service

import (
	"awesomeProject2/model"
	"fmt"
	"math/rand"
	"time"
)

// Player represents a playerAction in the card game.
type Player interface {
	ShowHand()
}

// Dealer represents a dealer in the card game.
type Dealer struct {
	Hand []model.Card
}

// HumanPlayer represents a human playerAction in the card game.
type HumanPlayer struct {
	Hand []model.Card
}

// ShowHand displays the hand of a playerAction.
func (p Dealer) ShowHand(isShow bool) {
	fmt.Printf("\nDealer's Hand (%d cards):\n", len(p.Hand))
	fmt.Printf("1: %s %s\n", p.Hand[0].Value, p.Hand[0].Suit)
	if isShow == false {
		fmt.Println("2: Face Down Card")
	} else {
		fmt.Printf("2: %s %s\n", p.Hand[1].Value, p.Hand[1].Suit)
	}
}

// ShowHand displays the hand of a playerAction.
func (p HumanPlayer) ShowHand() {
	fmt.Printf("\nPlayer's Hand (%d cards):\n", len(p.Hand))
	for i, card := range p.Hand {
		fmt.Printf("%d: %s %s\n", i+1, card.Value, card.Suit)
	}
}

// NewDeck generates a deckPlay.go of cards.
func NewDeck() []model.Card {
	suits := []string{"♥", "♦", "♣", "♠"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	deck := make([]model.Card, 0)

	for _, suit := range suits {
		for _, value := range values {
			card := model.Card{Suit: suit, Value: value}
			deck = append(deck, card)
		}
	}

	return deck
}

func DeleteExistedCardInDeck(deck []model.Card, totalCards int) []model.Card {
	if totalCards >= len(deck) {
		// Return an empty slice if the number of cards to delete is greater than or equal to the length of the deckPlay
		return []model.Card{}
	}
	return deck[totalCards:]
}

// ShuffleDeck shuffles a deckPlay.go of cards.
func ShuffleDeck(deck []model.Card) []model.Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

// DealCards deals cards to players.
func DealCards(deck []model.Card, numCards int) ([]model.Card, []model.Card) {
	dealerHand := make([]model.Card, 0)
	playerHand := make([]model.Card, 0)

	for i := 0; i < numCards; i++ {
		dealerHand = append(dealerHand, model.Card{
			Value:    deck[i].Value,
			Suit:     deck[i].Suit,
			IsDealer: true,
		})
		playerHand = append(playerHand, deck[i+1])
	}

	return dealerHand, playerHand
}
