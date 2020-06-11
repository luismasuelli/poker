package low

import (
	"github.com/luismasuelli/poker-go/assets/cards"
	"github.com/luismasuelli/poker-go/rules/french/std52/evaluators/card7"
	"github.com/luismasuelli/poker-go/rules/french/std52/evaluators/common"
)

// Computes the best power (and best cards combinations) of the given 7 cards.
// The power is taken considering the first combination of cards having the
// LOWER possible power (not considering also flushes and straights), using
// lowball rule for Ace.
// This supports any 7-card lowball game and mode, like:
// - Razz.
// - 7-Cards stud Hi/Lo when owning 7 cards (less than 8 active players on showdown).
// - 7-Cards stud Hi/Lo when owning 6 cards, and 1 in community (8 active players on
//   showdown).
func Power(hand []cards.Card, community []cards.Card) (best uint32, power uint64) {
	fullHand := common.AddCards(hand, community)
	power = ^uint64(0)
	best = 0
	for _, combination := range card7.Combinations {
		bits := combination[0]
		handBits, _ := common.Pick(fullHand, combination, common.LowballRanks)
		currentPower := common.Std52LowballPower(handBits)
		if currentPower < power {
			best = bits
			power = currentPower
		}
	}
	return
}
