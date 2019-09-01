//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

// Suit of the card
type Suit uint8

// Available suits of the cards
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

// Rank of the cards
type Rank uint8

// Available ranks of the cards
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card is the representation of each card in a deck
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}
