package service

import (
	"awesomeProject2/model"
	"fmt"
)

var deckPlay []model.Card

// StartRound simulates one round of the card game.
func StartRound() map[int]string {
	indexHandMapResult := make(map[int]string)
	// Check if deckPlay is nil or empty, and create a new one if needed
	deckPlay = NewDeck()
	ShuffleDeck(deckPlay)
	var dealer Dealer
	var player HumanPlayer

	dealer.Hand, player.Hand = DealCards(deckPlay, 2)
	//Draw card -> delete card in deckPlay.go
	deckPlay = DeleteExistedCardInDeck(deckPlay, len(dealer.Hand)+len(player.Hand))

	// Display hands of the dealer and playerAction
	dealer.ShowHand(false)
	player.ShowHand()

	// Calculate playerAction if score = 21 -> BLACK JACK -> PLAYER WIN
	playerBlackJack := CalculateScore(player.Hand)
	dealerBlackJack := CalculateScore(dealer.Hand)
	if playerBlackJack == 21 && dealerBlackJack == 21 {
		fmt.Println("----------------RESULT----------------")
		dealer.ShowHand(true)
		indexHandMapResult[1] = model.Draw
		return indexHandMapResult
	} else if playerBlackJack == 21 {
		fmt.Println("----------------RESULT----------------")
		indexHandMapResult[1] = model.PlayerWon + " " + model.BlackJack
		return indexHandMapResult
	} else if dealerBlackJack == 21 {
		fmt.Println("----------------RESULT----------------")
		dealer.ShowHand(true)
		indexHandMapResult[1] = model.DealerWon + " " + model.BlackJack
		return indexHandMapResult
	}

	// Add game logic for the round here
	playerScore := PlayerTurn(player.Hand)
	for index, score := range playerScore {
		if score > 21 {
			indexHandMapResult[index] = model.DealerWon
			return indexHandMapResult
		}
	}
	dealerScore := DealerTurn(dealer.Hand)
	for index, score := range playerScore {
		if score == dealerScore {
			indexHandMapResult[index] = model.Draw
		} else if dealerScore > score && dealerScore <= 21 || score > 21 {
			indexHandMapResult[index] = model.DealerWon
		} else if score > dealerScore && score <= 21 || dealerScore > 21 {
			indexHandMapResult[index] = model.PlayerWon
		}
		index++
	}

	return indexHandMapResult
}

func CalculateScore(hand []model.Card) int {
	cardValues := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"10": 10, "J": 10, "Q": 10, "K": 10,
		"A": 11,
	}

	score := 0
	aceCount := 0

	for _, card := range hand {
		if value, ok := cardValues[card.Value]; ok {
			if card.Value == "A" {
				aceCount++
			} else {
				score += value

			}
		}
	}

	for aceCount > 0 {
		if score <= 10 {
			score += 11
		} else {
			score += 1
		}
		aceCount--
	}

	return score
}

func PlayerTurn(hand []model.Card) map[int]int {
	indexMapScore := make(map[int]int)
	for {
		var playerChoice int
		fmt.Println("-----------------ACTION-----------------")
		fmt.Println("1: Stand - 2: Draw - 3: Split . Choose an action (1-4): ")
		fmt.Scanln(&playerChoice)
		switch playerChoice {
		case 1:
			return PlayerStands(hand, indexMapScore, 1)
		case 2:
			return PlayerDraw(hand, indexMapScore, 1)
		case 3:
			result := PlayerSplit(hand, indexMapScore)
			if result != nil {
				return result
			} else {
				fmt.Println("Invalid input for splitting. Please try again.")
			}
		default:
			fmt.Println("Invalid choice. Please enter a valid number.")
		}

	}
}

func DealerTurn(hands []model.Card) int {
	tempScore := CalculateScore(hands)
	fmt.Printf("\nDealer's Show hands (%d cards):\n", len(hands))
	if tempScore >= 17 {
		for i, card := range hands {
			fmt.Printf("%d: %s %s\n", i+1, card.Value, card.Suit)

		}
		fmt.Printf("Dealer score: %d\n", tempScore)
		return tempScore
	} else {
		score := 0
		for {
			newCard := DrawCard()
			fmt.Printf("Dealer draw card: %s %s\n", newCard.Value, newCard.Suit)

			hands = append(hands, newCard)
			score = CalculateScore(hands)
			if score >= 17 {
				for i, card := range hands {
					fmt.Printf("%d: %s %s\n", i+1, card.Value, card.Suit)

				}
				fmt.Printf("Dealer score: %d\n", score)
				return score
			}
		}
	}
}

func DrawCard() model.Card {
	card := deckPlay[0]
	deckPlay = DeleteExistedCardInDeck(deckPlay, 1)
	return card
}
