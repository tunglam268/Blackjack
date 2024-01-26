package service

import (
	"awesomeProject2/model"
	"fmt"
)

var playerAction HumanPlayer

func PlayerStands(hand []model.Card, indexMapScore map[int]int, index int) map[int]int {
	fmt.Println("Player stands.")
	playerAction.Hand = hand
	totalScore := CalculateScore(playerAction.Hand)
	playerAction.ShowHand()

	// Update playerAction's hand in the indexMapScore map
	indexMapScore[index] = totalScore

	fmt.Printf("Player score: %d\n", totalScore)
	return indexMapScore
}

func PlayerDraw(hand []model.Card, indexMapScore map[int]int, index int) map[int]int {
	//Draw 1 card && calculate score
	fmt.Println("Player draw")
	newCard := DrawCard()
	hand = append(hand, newCard)
	fmt.Printf("Player draw card: %s %s\n", newCard.Value, newCard.Suit)

	//Show hand after draw
	playerAction.Hand = hand
	playerAction.ShowHand()

	totalScore := CalculateScore(hand)
	fmt.Printf("Player score: %d\n", totalScore)
	//Check score > 21 or = 21 => still return
	if totalScore > 21 || totalScore == 21 {
		if totalScore > 21 {
			fmt.Println("Bust! Player draws cards totaling over 21.")
		}

		indexMapScore[index] = totalScore
		return indexMapScore
	}
	//Action draw && stand
	for {
		var drawAgain int
		fmt.Print("1:Yes - 2:No . Do you want to draw again? (yes/no): ")
		fmt.Scanln(&drawAgain)
		if drawAgain == 2 {
			//Stand no more draw
			return PlayerStands(playerAction.Hand, indexMapScore, 0)
		} else {
			newCard := DrawCard()
			hand = append(hand, newCard)
			//Show hand after draw
			playerAction.Hand = hand
			playerAction.ShowHand()

			//Calculate after draw
			fmt.Printf("Player draw card: %s %s\n", newCard.Value, newCard.Suit)
			totalScore := CalculateScore(hand)
			fmt.Printf("Player score: %d\n", totalScore)
			if totalScore > 21 || totalScore == 21 {
				if totalScore > 21 {
					fmt.Println("Bust! Player draws cards totaling over 21.")
				}

				indexMapScore[index] = totalScore
				return indexMapScore
			}
		}
	}
}

func PlayerSplit(hand []model.Card, indexMapScore map[int]int) map[int]int {
	if hand[0].Value != hand[1].Value {
		fmt.Print("The cards are not the same ")
		return nil
	}
	for i := 1; i <= 2; i++ {
		newCard := DrawCard()
		splitHand := append([]model.Card{hand[i-1]}, newCard)
		totalScore := CalculateScore(splitHand)

		playerAction.Hand = splitHand
		playerAction.ShowHand()
		fmt.Printf("Player score of hand %d split: %d\n", i, totalScore)

		if totalScore == 21 {
			indexMapScore[i] = totalScore
			break
		}

		var playerChoice int
		fmt.Println("-----------------ACTION-----------------")
		fmt.Println("1: Stand - 2: Draw - 3: Split . Choose an action (1-4): ")
		fmt.Scanln(&playerChoice)
		switch playerChoice {
		case 1:
			indexMapScore = PlayerStands(splitHand, indexMapScore, i)
		case 2:
			indexMapScore = PlayerDraw(splitHand, indexMapScore, i)
		default:
			fmt.Println("Invalid choice. Please enter a valid number.")
		}
	}

	return indexMapScore
}
