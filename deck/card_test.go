package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Heart, Rank: Ace})
	fmt.Println(Card{Suit: Diamond, Rank: Three})
	fmt.Println(Card{Suit: Spade, Rank: King})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Three of Diamonds
	// King of Spades
	// Joker
}

func TestNew(t *testing.T) {
	deck := New()

	expected := 13 * 4
	got := len(deck)
	if expected != got {
		t.Errorf("Wrong number of cards in new deck. Expected: %d - Got: %d.", expected, got)
	}

	deck = New(DefaultShuffle)
	expected = 13 * 4
	got = len(deck)
	if expected != got {
		t.Errorf("Wrong number of cards in new deck. Expected: %d - Got: %d.", expected, got)
	}

	// Naive implementation for initial testing, this will fail once in 52 tries
	if (deck[0] == Card{Suit: Spade, Rank: Ace}) {
		t.Errorf("Wrong first card: %s", deck[0])
	}

	deck = New(DefaultSort)
	expected = 13 * 4
	got = len(deck)
	if expected != got {
		t.Errorf("Wrong number of cards in new deck. Expected: %d - Got: %d.", expected, got)
	}

	if (deck[0] != Card{Suit: Spade, Rank: Ace}) {
		t.Errorf("Wrong first card: %s", deck[0])
	}
}
