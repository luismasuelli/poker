package badugi

import (
	"github.com/luismasuelli/poker/assets/cards"
	. "github.com/luismasuelli/poker/assets/cards/french"
	"testing"
)

func testBadugi(t *testing.T, expectedBest uint32, expectedPower uint64, cards ...cards.Card) {
	best, power := Power(cards, nil)
	if best != expectedBest || power != expectedPower {
		t.Errorf(
			"Got best: %#04b vs. expected: %#04b, power: %#021b vs. expected: %#021b",
			best, expectedBest, power, expectedPower,
		)
	}
}

func TestNaiveCases(t *testing.T) {
	testBadugi(t, 0b1111, 0b000000000001111000000, CA, H2, D3, S4)
	testBadugi(t, 0b0111, 0b010000000000111000000, CA, H2, D3, S3)
	testBadugi(t, 0b0101, 0b101100000000000000000, CK, HK, HQ, SQ)
	testBadugi(t, 0b1010, 0b100100000000001000000, CK, CQ, CA, SA)
	testBadugi(t, 0b1100, 0b100100000000001000000, CK, CQ, CA, SQ)
	testBadugi(t, 0b1100, 0b101000000000001000000, CK, CQ, CA, SK)
	testBadugi(t, 0b1001, 0b101100000000000000000, CK, CQ, SK, SQ)
}
