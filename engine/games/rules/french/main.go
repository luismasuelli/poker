package french

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	"github.com/luismasuelli/poker-go/engine/games/cards/french"
)

// A standard french deck contains a sequence of standard french
// cards, and will act on it, change it, and so.
type Deck struct {
	cards []french.Card
}

// Creates a deck with the given french cards.
func NewDeck(cards ...french.Card) *Deck {
	return &Deck{
		cards: cards,
	}
}

// The length of a deck is the length of the underlying array.
func (deck *Deck) Len() int {
	return len(deck.cards)
}

// Returns a copy of the current deck.
func (deck *Deck) Copy() cards.Deck {
	currentCards := make([]french.Card, len(deck.cards))
	copy(currentCards, deck.cards)
	return &Deck{cards: currentCards}
}

// Swaps two cards inside the deck.
func (deck *Deck) Swap(i, j int) {
	deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
}

// Deals n cards from the top of the deck.
func (deck *Deck) Deal(n int) []cards.Card {
	peeked := deck.Peek(n)
	deck.cards = deck.cards[0 : len(deck.cards)-len(peeked)]
	return peeked
}

// Peeks n cards from the top of the deck (a non-destructive
// way of dealing cards).
func (deck *Deck) Peek(n int) []cards.Card {
	baseLength := len(deck.cards)

	if n < 1 {
		panic(cards.ErrDealBadCount)
	} else if n > baseLength {
		panic(cards.ErrDealNotEnough)
	} else {
		newLength := baseLength - n
		source := deck.cards[newLength:baseLength]
		result := make([]cards.Card, len(source))
		for index := 0; index < n; index++ {
			result[index] = source[n-1-index]
		}
		return result
	}
}

// Stacks new cards onto the deck, in the order
// they are specified.
func (deck *Deck) Stack(cards []cards.Card) {
	newLength := len(cards)
	baseLength := len(deck.cards)

	if newLength == 0 {
		return
	}

	newCards := make([]french.Card, newLength+baseLength)
	copy(newCards, deck.cards)
	for index := 0; index < newLength; index++ {
		newCards[baseLength+index] = cards[index].(french.Card)
	}

	deck.cards = newCards
}

// Queues new cards under the deck, in the order
// they are specified.
func (deck *Deck) Queue(cards []cards.Card) {
	newLength := len(cards)
	baseLength := len(deck.cards)

	if newLength == 0 {
		return
	}

	newCards := make([]french.Card, newLength+baseLength)
	for index := 0; index < newLength; index++ {
		newCards[newLength-1-index] = cards[index].(french.Card)
	}
	copy(newCards[newLength:newLength+baseLength], deck.cards)

	deck.cards = newCards
}
