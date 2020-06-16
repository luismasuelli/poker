package high

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	"github.com/luismasuelli/poker-go/engine/games/rules/french/std52/evaluators/card7"
	"github.com/luismasuelli/poker-go/engine/games/rules/french/std52/evaluators/common"
)

// Computes the best power (and best cards combinations) of the given 7 cards.
// The power is taken considering the first combination of cards having the
// GREATER possible power (considering also flushes and straights).
// This supports any 7-card high game and mode, like:
// - Texas Hold'Em.
// - 7-Cards stud when owning 7 cards (less than 8 active players on showdown).
// - 7-Cards stud when owning 6 cards, and 1 in community (8 active players on
//   showdown).
func Power(hand []cards.Card, community []cards.Card) (best uint32, power uint64) {
	fullHand := common.AddCards(hand, community)
	power = 0
	best = 0
	for _, combination := range card7.Combinations {
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
