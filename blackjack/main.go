package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dlokkers/gophercises/deck"
)

// Hand represents the players or the dealers hand
type Hand []deck.Card

func main() {
	cards := deck.New(deck.DefaultShuffle)

	var playerHand Hand
	var dealerHand Hand

	for i := 0; i < 2; i++ {
		playerHand = append(playerHand, cards.Draw(deck.Up))
		dealerHand = append(dealerHand, cards.Draw(deck.Down))
	}

	dealerHand[1].Flip()

	fmt.Println("Dealer:")
	fmt.Println(dealerHand)

	fmt.Println("Player:")
	fmt.Println(playerHand)

	if playerHand.value() < 21 || dealerHand.value() < 21 {
		playerHand, cards = playerTurn(playerHand, cards)
		if playerHand.value() > 0 {
			dealerHand, cards = dealerTurn(dealerHand, cards)
		} else {
			dealerHand[0].Flip()
		}
	} else {
		dealerHand[0].Flip()
	}

	fmt.Println("-------------------------")
	fmt.Print("Dealer: ")
	fmt.Println(dealerHand.value())
	fmt.Println(dealerHand)

	fmt.Print("Player: ")
	fmt.Println(playerHand.value())
	fmt.Println(playerHand)

	if playerHand.value() > dealerHand.value() {
		fmt.Println("Player wins")
	} else if dealerHand.value() > playerHand.value() {
		fmt.Println("Dealer wins")
	} else {
		fmt.Println("Draw!")
	}
}

func playerTurn(h Hand, d deck.Deck) (Hand, deck.Deck) {
	var call string

	for call != "stand" {
		call = askCall()

		if call == "hit" {
			h = append(h, d.Draw(deck.Up))
			fmt.Println(h)

			if h.value() == 0 {
				return h, d
			}
		}
	}

	return h, d
}

func dealerTurn(h Hand, d deck.Deck) (Hand, deck.Deck) {
	fmt.Println("Dealers Turn")
	h[0].Flip()

	for h.value() < 17 {
		h = append(h, d.Draw(deck.Up))
		fmt.Println(h)

		if h.value() == 0 {
			return h, d
		}
	}

	return h, d
}

func askCall() string {
	fmt.Print("Hit or Stand?: ")

	callIO := bufio.NewScanner(os.Stdin)
	callIO.Scan()
	return strings.ToLower(callIO.Text())
}

func (h Hand) String() string {
	hand := ""
	for _, c := range h {
		hand = hand + fmt.Sprintln(c)
	}

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

	if total > 21 {
		total = 0
	}

	return total
}

func min(x, y uint8) uint8 {
	if x < y {
		return x
	}
	return y
}
