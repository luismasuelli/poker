package low

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	. "github.com/luismasuelli/poker-go/engine/games/cards/french"
	"testing"
)

func testHandPower(t *testing.T, expectedPower uint64, expectedBest uint32, cards ...cards.Card) {
	best, power := Power(cards, nil)
	if power != expectedPower || best != expectedBest {
		t.Errorf(
			"Testing hand: %v\nexpected power: %#064b\n     got power: %#064b\nexpected best: %#07b\n     got best: %#07b\n",
			cards, expectedPower, power, expectedBest, best,
		)
	}
}

func TestHandPowers(t *testing.T) {
	// 4 of a kind: are easily avoided in 7-Cards low.
	// Full house: will only occur in an XXXXYYY scenario, avoiding XXXXY.
	testHandPower(t, 0b0110000000000000000000000100000100000000000, 0b11111, SQ, S5, C5, HQ, D5, H5, DQ)
	testHandPower(t, 0b0110000000000000000000000000011000000000000, 0b101111, SA, SK, CK, HA, DK, DA, HK)
	testHandPower(t, 0b0110000000000000000000100000001000000000000, 0b11111, H8, SK, CK, H8, D8, HK, DK)
	// 3 of a kind: are easily avoided in 7-Cards low.
	// Double Pair.
	testHandPower(t, 0b0010000000000000000010100000000100000000000, 0b1011011, CT, ST, DT, D8, C8, H8, CQ)
	testHandPower(t, 0b0010000000000000000000010000010001000000000, 0b0101111, DA, SA, HT, D7, CA, C7, H7)
	testHandPower(t, 0b0010000000000000000000000010100000100000000, 0b11111, C2, D2, S4, H4, S9, H2, D4)
	// Pair.
	testHandPower(t, 0b0001000000000000000010000000000100010001000, 0b11111, CT, ST, D4, C8, HQ, DQ, CQ)
	testHandPower(t, 0b0001000000000000000000000000010001001000100, 0b110111, CT, DA, SA, HT, C3, H7, C7)
	testHandPower(t, 0b0001000000000000000000000000100000100001100, 0b11111, C2, D2, S3, H4, S9, S4, H9)
	// Bust / High Card.
	testHandPower(t, 0b0000000000000000000000000000000000111101000, 0b11111, S7, S9, C6, H4, S8, D7, H7)
	testHandPower(t, 0b0000000000000000000000000000000001010100101, 0b1011011, HA, D6, H6, D8, D3, H8, HT)
	testHandPower(t, 0b0000000000000000000000000000000001001001101, 0b1111001, DT, HT, ST, CA, S3, H4, D7)
}
