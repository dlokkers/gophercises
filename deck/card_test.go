package deck

import "fmt"

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
