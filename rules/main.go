package rules

import "github.com/luismasuelli/poker/assets/cards"

// Decks keep collections of cards. They can copy
// themselves, deal one or more cards (panicking
// when the number of cards to deal is < 1 or when
// there are no left cards), and provide a swapper
// function to use in a shuffle. Also: deck will
// hav their own set of rules (this means: they
// may choose to accept only a subset of the
// valid cards in the set they belong to, e.g.
// a 52-cards deck will allow all the french cards
// but the wildcards).
type Deck interface {
	// Copies the entire deck, creating a new one.
	// This method should only be run on templates,
	// and not expected to be safe on concurrency.
	Copy() Deck
	// Gets the length of a deck (considering its
	// remaining cards).
	Len() int
	// A swapper function to use in a shuffler.
	// Panics should be normally raised in this
	// functions when indices are out of bounds.
	Swap(i, j int)
	// Deals cards from the top of the deck (i.e.
	// it unstacks the last cards). The cards are
	// returned in unstacked order (i.e. reverse
	// of the stack order).
	Deal(int) []cards.Card
	// Peeks cards from the top of the deck (i.e.
	// it unstacks the last cards). The cards are
	// returned in unstacked order (i.e. reverse
	// of the stack order). The difference with
	// Deal(int) is that this method does not
	// actually remove the cards, but anticipates
	// or simulates it.
	Peek(int) []cards.Card
	// Returns cards to the top of the deck (i.e.
	// it stacks them back onto the deck), one by
	// one in the order they are given. It must
	// panic if at least one card is of unexpected
	// type, or nil. It must be a no-op on empty
	// or nil array.
	Stack([]cards.Card)
	// Returns cards to the bottom of the deck
	// (i.e. it queues them back to the bottom
	// of the deck), one by one in the order
	// they are given. It must panic if at least
	// one card is of unexpected type, or nil. It
	// must be a no-op on empty or nil array.
	Queue([]cards.Card)
}
