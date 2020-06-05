package cards

import "errors"

type CardFlags uint8

const (
	// This is a valid card. It will be properly evaluated.
	Valid = CardFlags(0)

	// This is an invalid card, according to the implementation's
	// cards set. Trying to evaluate a card with this flag should
	// panic in the evaluator.
	Invalid = CardFlags(4)

	// This is a hidden card. It is an invalid card for the purpose
	// of evaluation, and it should display a card's back on client
	// side.
	Unknown = Invalid | CardFlags(1)

	// This is an unexpected card. It is an invalid card for the
	// purpose of evaluation, and if by chance it reaches the
	// client side, it should display a wrong card (a special
	// asset image or message).
	Unexpected = Invalid | CardFlags(2)
)

// Raised by panic when attempting to deal < 1 cards.
var ErrDealBadCount = errors.New("invalid number of cards to deal")

// Raised by panic when attempting to deal > length(deck) cards.
var ErrDealNotEnough = errors.New("not enough cards to deal")

// Cards are nothing by themselves, save for their
// visual representation. This is done in their
// "face" method. The "face" should be matched in
// front-end with an appropriate representation.
//
// Cards can also come in different statuses like:
// a valid (evaluable) card, an unknown card, or an
// invalid (unexpected) card.
//
// Each implementation will define the meaning and
// components of their values by themselves.
type Card interface {
	// Face string of this card.
	Face() string
	// Flags of this card, telling whether it is
	// valid, unknown, or unexpected.
	Flags() CardFlags
}

// Decks keep collections of cards. They can copy
// themselves, deal one or more cards (panicking
// when the number of cards to deal is < 1 or when
// there are no left cards), and provide a swapper
// function to use in a shuffle.
type Deck interface {
	// Copies the entire deck, creating a new one.
	Copy() Deck
	// Gets the length of a deck (considering its
	// remaining cards).
	Len() int
	// A swapper function to use in a shuffler.
	Swapper() func(i, j int)
	// Deals cards from the top of the deck (i.e.
	// it unstacks the last cards). The cards are
	// returned in unstacked order (i.e. reverse
	// of the stack order).
	Deal(int) []Card
	// Peeks cards from the top of the deck (i.e.
	// it unstacks the last cards). The cards are
	// returned in unstacked order (i.e. reverse
	// of the stack order). The difference with
	// Deal(int) is that this method does not
	// actually remove the cards, but anticipates
	// or simulates it.
	Peek(int) []Card
	// Returns cards to the top of the deck (i.e.
	// it stacks them back onto the deck), one by
	// one in the order they are given. It must
	// panic if at least one card is of unexpected
	// type, or nil. It must be a no-op on empty
	// or nil array.
	Stack([]Card)
	// Returns cards to the bottom of the deck
	// (i.e. it queues them back to the bottom
	// of the deck), one by one in the order
	// they are given. It must panic if at least
	// one card is of unexpected type, or nil. It
	// must be a no-op on empty or nil array.
	Queue([]Card)
}