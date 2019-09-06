//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

// Suit of the card
type Suit uint8

// Suits available
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank of the cards
type Rank uint8

// Ranks available
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

var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Card is the representation of each card in a deck
type Card struct {
	Suit
	Rank
}

// Cards represents a deck or cards+
type Cards []Card

func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// New creates a new deck of cards, including set of optional
// parameters
func New() []Card {
	var cards []Card

	for _, s := range suits {
		for _, r := range ranks {
			cards = append(cards, Card{s, r})
		}
	}

	return cards
}

// Len defines the size of the deck, to implement Sort
func (c Cards) Len() int {
	return len(c)
}

// Swap defines a means to swap cards for sorting purposes
func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Less defines an ordering in cards for sorting purposes
func (c Cards) Less(i, j int) bool {
	return int(c[i].Suit)*len(ranks)+int(c[j].Rank) <
		int(c[j].Suit)*len(ranks)+int(c[j].Rank)
}
