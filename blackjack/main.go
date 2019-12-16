package main

import (
	"fmt"

	"github.com/dlokkers/gophercises/deck"
)

// Hand represents the players or the dealers hand
type Hand []deck.Card

func main() {
	cards := deck.New(deck.DefaultShuffle)

	var playerHand Hand
	var dealerHand Hand

	for i := 0; i < 2; i++ {
		playerHand = append(playerHand, cards.Draw())
		dealerHand = append(dealerHand, cards.Draw())
	}

	fmt.Println(dealerHand)
	fmt.Println(playerHand)

	if playerHand.value() > dealerHand.value() {
		fmt.Println("Player wins")
	} else if dealerHand.value() > playerHand.value() {
		fmt.Println("Dealer wins")
	} else {
		fmt.Println("Draw!")
	}
}

func (h Hand) String() string {
	hand := ""
	for _, c := range h {
		hand = hand + fmt.Sprintln(c)
	}
	hand = hand + fmt.Sprintf("Total points: %v\n", h.value())

	return hand
}

// value returns the total value of the cards in hand
func (h Hand) value() uint8 {
	var total uint8
	aces := 0

	for _, c := range h {
		if c.Rank == deck.Ace {
			aces++
		}
		total += min(10, uint8(c.Rank))
	}

	if aces > 0 {
		if total+10 <= 21 {
			total += 10
		}
	}

	return total
}

func min(x, y uint8) uint8 {
	if x < y {
		return x
	}
	return y
}
