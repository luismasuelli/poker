package rankers

import (
	"github.com/luismasuelli/poker/assets/cards"
	"errors"
)

// Panicked when a ranker is trying to rank a nil value.
var ErrNilCard = errors.New("nil card being ranked")

// Panicked when a ranker is trying to rank an invalid card.
var ErrInvalidCard = errors.New("invalid card being ranked")

// Ranks a card according to an internal rule set.
// If the card is not valid, or the card is of a
// different type, a panic will be raised.
type Ranker struct {
	scale func(cards.Card) uint8
}

// Creates a new ranker, using a custom scaling function.
func NewRanker(scale func(cards.Card) uint8) Ranker {
	return Ranker{scale}
}

// Attempts to rank a card, using the given scaling function.
func (ranker Ranker) Rank(card cards.Card) uint8 {
	if card == nil {
		panic(ErrNilCard)
	} else if card.Flags() & cards.Invalid > 0 {
		panic(ErrInvalidCard)
	} else {
		return ranker.scale(card)
	}
}