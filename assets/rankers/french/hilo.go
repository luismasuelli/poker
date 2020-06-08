package french

import (
	"github.com/luismasuelli/poker-go/assets/cards"
	"github.com/luismasuelli/poker-go/assets/cards/french"
	"github.com/luismasuelli/poker-go/assets/rankers"
)

// Ranks a french card, all of them but wildcards
// (they are forbidden in this ranking), with A-high.
func HighRanker(card cards.Card) uint8 {
	frenchCard := card.(french.Card)
	if frenchCard.Suit() == french.Wildcard {
		panic(rankers.ErrInvalidCard)
	}
	return frenchCard.Rank() - 2
}

// Ranks a french card, all of them but wildcards
// (they are forbidden in this ranking), with A-low.
func LowRanker(card cards.Card) uint8 {
	frenchCard := card.(french.Card)
	if frenchCard.Suit() == french.Wildcard {
		panic(rankers.ErrInvalidCard)
	} else if frenchCard.Rank() == 14 {
		return 0
	} else {
		return frenchCard.Rank() - 1
	}
}

// Ranks a french card with A-high. Wildcards count
// as A, and are ranked accordingly.
func Caribbean(card cards.Card) uint8 {
	frenchCard := card.(french.Card)
	if frenchCard.Suit() == french.Wildcard || frenchCard.Rank() == 14 {
		return 0
	} else {
		return frenchCard.Rank() - 1
	}
}