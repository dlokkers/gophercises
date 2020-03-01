//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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
	Face
}

// Face describes if a card is facing up or down
type Face uint8

// Describing the states a card can be in
const (
	Up Face = iota
	Down
)

// Deck represents a deck of multiple cards
type Deck []Card

func (c Card) String() string {
	if c.Face == Down {
		return "******"
	}
	if c.Suit == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// New creates a new deck of cards, including set of optional
// parameters
func New(opts ...func(Deck) Deck) Deck {
	var cards Deck

	for _, s := range suits {
		for _, r := range ranks {
			cards = append(cards, Card{s, r, Down})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// Flip flips the card over
func (c *Card) Flip() {
	if c.Face == Up {
		c.Face = Down
	} else {
		c.Face = Up
	}
}

// Len defines the size of the deck, to implement Sort
func (d Deck) Len() int {
	return len(d)
}

// Swap defines a means to swap cards for sorting purposes
func (d Deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less defines an ordering in cards for sorting purposes
func (d Deck) Less(i, j int) bool {
	return int(d[i].Suit)*len(ranks)+int(d[j].Rank) <
		int(d[j].Suit)*len(ranks)+int(d[j].Rank)
}

// DefaultSort implements a default sorting for the deck
func DefaultSort(d Deck) Deck {
	sort.Sort(d)
	return d
}

// DefaultShuffle offers a build in shuffle method to use on the deck
func DefaultShuffle(d Deck) Deck {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make(Deck, len(d))
	perm := r.Perm(len(d))

	for i, randIndex := range perm {
		ret[i] = d[randIndex]
	}
	return ret
}

// AddJokers adds a joker to the Deck
func AddJokers(n int) func(Deck) Deck {
	return func(d Deck) Deck {
		for i := 0; i < n; i++ {
			d = append(d, Card{Suit: Joker})
		}
		return d
	}
}

// Draw returns and removes the first card in the deck
func (d *Deck) Draw(f Face) Card {
	var c Card

	c, *d = (*d)[0], (*d)[1:]
	c.Face = f

	return c
}
