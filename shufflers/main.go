package shufflers

import "github.com/luismasuelli/poker-go/rules"

// Shufflers provide only one method to shuffle any
// given deck (regardless the deck type). The
// shuffling is always in-place, and it will make
// use of the .Swap(i, j) and .Len() functions
// inside a deck.
type Shuffler interface {
	Shuffle(deck rules.Deck)
}
