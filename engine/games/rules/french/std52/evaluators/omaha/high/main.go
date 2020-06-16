package high

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	"github.com/luismasuelli/poker-go/engine/games/rules/french/std52/evaluators/common"
	"github.com/luismasuelli/poker-go/engine/games/rules/french/std52/evaluators/omaha"
)

// Computes the best power (and best cards combinations) of the given 9 cards.
// The power is taken considering the first combination of cards having the
// GREATER possible power (considering also flushes and straights).
// This supports Omaha high (or the high part in hi/lo games).
func Power(hand []cards.Card, community []cards.Card) (best uint32, power uint64) {
	fullHand := common.AddCards(hand, community)
	power = 0
	best = 0
	for _, combination := range omaha.Combinations {
		bits := combination[0]
		handBits, suitBits := common.Pick(fullHand, combination, common.HighRanks)
		currentPower := common.Std52HighPower(handBits, suitBits != 0)
		if currentPower > power {
			best = bits
			power = currentPower
		}
	}
	return
}
