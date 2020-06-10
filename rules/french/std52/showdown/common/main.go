package common

var HighRanges = []uint64{
	0, // Code 0 does not exist.
	0, // Code 1 does not exist.
	0b000000000000000000000000000000000000001, // 2,...
	0b000000000000000000000000000000000001000,
	0b000000000000000000000000000000001000000,
	0b000000000000000000000000000001000000000,
	0b000000000000000000000000001000000000000,
	0b000000000000000000000001000000000000000,
	0b000000000000000000001000000000000000000,
	0b000000000000000001000000000000000000000,
	0b000000000000001000000000000000000000000, // 10
	0b000000000001000000000000000000000000000, // J
	0b000000001000000000000000000000000000000, // Q
	0b000001000000000000000000000000000000000, // K
	0b001000000000000000000000000000000000000, // A
	0, // Code 15 does not exist.
}

var LowballRanges = []uint64{
	0, // Code 0 does not exist.
	0, // Code 1 does not exist.
	0b000000000000000000000000000000000001000, // 2,...
	0b000000000000000000000000000000001000000,
	0b000000000000000000000000000001000000000,
	0b000000000000000000000000001000000000000,
	0b000000000000000000000001000000000000000,
	0b000000000000000000001000000000000000000,
	0b000000000000000001000000000000000000000,
	0b000000000000001000000000000000000000000,
	0b000000000001000000000000000000000000000, // 10
	0b000000001000000000000000000000000000000, // J
	0b000001000000000000000000000000000000000, // Q
	0b001000000000000000000000000000000000000, // K
	0b000000000000000000000000000000000000001, // A
	0, // Code 15 does not exist.
}

// Tries detecting a straight in the hand bits. Returns
// a 0-power if no straight was detected. This function
// is not meant to be used in lowball.
func std52TestStraight(handBits uint64) uint64 {
	const straight5 = 0b001000000000000000000000000001001001001
	var straightSeq uint64 = 0b1001001001001

	if handBits&straight5 == straight5 {
		return 0b0000000001000
	}
	for i := 0; i < 9; i++ {
		if straightSeq&handBits == straightSeq {
			// E.g. for i=0, the range is 6, which has a
			// bit shift of 4.
			return 1 << (4 + i)
		}
		// This moves the bitmask one range ahead.
		straightSeq <<= 3
	}
	return 0
}

// Tries detecting all the patterns in 5 cards, as follows:
// - 4 of a kind.
// - 3 of a kind.
// - 2 of a kind (first and second).
// - Kicker (first, second, and third).
// Each of these patterns will be 1-hot rank vectors (not
// shifted - that will be done later) in their lower 13
// bits, and having a number 0..13 in the upper 4 bits
// telling how many patterns of this type.
//
// Patterns are useful both in standard and lowball, but
// the bits will have different meanings in either modes.
func std52TestPatterns(handBits uint64) (oak4 uint64, oak3 uint64, oak2 uint64, kicker uint64) {
	var handBitsToShift = handBits
	var activate uint64 = 1
	const add1 = 1 << 60
	for i := 0; i < 39; i += 3 {
		bits := handBitsToShift & 7
		switch bits {
		case 4:
			oak4 |= activate
			oak4 += add1
		case 3:
			oak3 |= activate
			oak3 += add1
		case 2:
			oak2 |= activate
			oak2 += add1
		case 1:
			kicker |= activate
			kicker += add1
		case 0:
			// Nothing - just move further.
		default:
			// It is invalid to have 5+ equal cards in standard 52.
			return 0, 0, 0, 0
		}
		handBits &= ^(bits << i)
		handBitsToShift >>= 3
		activate <<= 1
	}
	return
}

// Assumes a hand of 5 cards out of 52, and evaluates the
// combinations: Straight-Flush, 4oak, Full House, Flush,
// Straight, 3oak, Double Pair, Pair, Bust. In this mode,
// the Ace counts high (save for the low straight).
//
// It receives two arguments:
// - A set of bits AAAKKKQQQJJJTTT999888777666555444333222.
//   Each XXX can be 000 to 100 (100 meaning 4 of a kind: X).
// - A flag telling whether they all have the same suit,
//   regardless the particular suit.
//
// The returned power comes in different flavors:
// - Straight Flush: [1000][00000000000000000000000000][vvvvvvvvvvvvv]
//   with vvvvvvvvvvvvv a 1-hot rank vector telling the reach of the
//   straight.
// - 4 of a kind: [0111][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 4 equal
//   cards, and kkk... a 1-hot rank vector telling the rank of the
//   kicker.
// - Full House: [0110][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 3 equal
//   cards, and kkk... a 1-hot rank vector telling the rank of the
//   2 equal cards.
// - Flush: [0101][00000000000000000000000000][vvvvvvvvvvvvv]
//   with vvvvvvvvvvvvv a combination of 1-hot rank vectors telling the
//   rank of each involved card.
// - Straight: [0100][00000000000000000000000000][vvvvvvvvvvvvv]
//   with vvvvvvvvvvvvv a 1-hot rank vector telling the reach of the
//   straight.
// - 3 of a kind: [0011][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 3 equal
//   cards, and kkk... the combination of two 1-hot rank vectors
//   telling the rank of the kickers.
// - Double Pair: [0010][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a combination of two 1-hot rank vectors telling the
//   ranks of the 2/2 equal cards, and kkk... a 1-hot rank vector
//   telling the rank of the kicker.
// - Pair: [0001][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 2 equal
//   cards, and kkk... a combination of 3 1-hot rank vectors telling
//   the ranks of the kickers.
// - Bust (High Cards): [0000][00000000000000000000000000][vvvvvvvvvvvvv]
//   with vvvvvvvvvvvvv a combination of 1-hot rank vector telling the
//   rank of each involved card.
func Std52HighPower(handBits uint64, hasFlush bool) uint64 {
	if hasFlush {
		// Straight-flush or flush.
		if power := std52TestStraight(handBits); power != 0 {
			return 8<<39 | power
		} else {
			return 5<<39 | handBits
		}
	} else if power := std52TestStraight(handBits); power != 0 {
		return 4<<39 | power
	} else {
		const mask = (1 << 13) - 1
		const twoPairs = 1 << 61
		oak4, oak3, oak2, kicker := std52TestPatterns(handBits)
		if oak4 != 0 {
			return 7<<39 | (oak4&mask)<<13 | (kicker & mask)
		} else if oak3 != 0 {
			if oak2 != 0 {
				return 6<<39 | (oak3&mask)<<13 | (oak2 & mask)
			} else {
				return 3<<39 | (oak3&mask)<<13 | (kicker & mask)
			}
		} else if oak2 != 0 {
			if oak2 > twoPairs {
				return 2<<39 | (oak2&mask)<<13 | (kicker & mask)
			} else {
				return 1<<39 | (oak2&mask)<<13 | (kicker & mask)
			}
		} else {
			return kicker
		}
	}
}

// Assumes a hand of 5 cards out of 52, and evaluates the
// combinations: 4oak, Full House, 3oak, Double Pair, Pair,
// Bust. In this mode, the Ace counts always low.
//
// It receives two arguments:
// - A set of bits KKKQQQJJJTTT999888777666555444333222AAA.
//   Each XXX can be 000 to 100 (100 meaning 4 of a kind: X).
//
// The returned power comes in different flavors:
// - 4 of a kind: [0111][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 4 equal
//   cards, and kkk... a 1-hot rank vector telling the rank of the
//   kicker.
// - Full House: [0110][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 3 equal
//   cards, and kkk... a 1-hot rank vector telling the rank of the
//   2 equal cards.
// - 3 of a kind: [0011][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 3 equal
//   cards, and kkk... the combination of two 1-hot rank vectors
//   telling the rank of the kickers.
// - Double Pair: [0010][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a combination of two 1-hot rank vectors telling the
//   ranks of the 2/2 equal cards, and kkk... a 1-hot rank vector
//   telling the rank of the kicker.
// - Pair: [0001][0000000000000][rrrrrrrrrrrrr][kkkkkkkkkkkkk]
//   with rrr... a 1-hot rank vector telling the rank of the 2 equal
//   cards, and kkk... a combination of 3 1-hot rank vectors telling
//   the ranks of the kickers.
// - Bust (High Cards): [0000][00000000000000000000000000][vvvvvvvvvvvvv]
//   with vvvvvvvvvvvvv a combination of 1-hot rank vector telling the
//   rank of each involved card.
func Std52LowballPower(handBits uint64) uint64 {
	const mask = (1 << 13) - 1
	const twoPairs = 1 << 61
	oak4, oak3, oak2, kicker := std52TestPatterns(handBits)
	if oak4 != 0 {
		return 7<<39 | (oak4&mask)<<13 | (kicker & mask)
	} else if oak3 != 0 {
		if oak2 != 0 {
			return 6<<39 | (oak3&mask)<<13 | (oak2 & mask)
		} else {
			return 3<<39 | (oak3&mask)<<13 | (kicker & mask)
		}
	} else if oak2 != 0 {
		if oak2 > twoPairs {
			return 2<<39 | (oak2&mask)<<13 | (kicker & mask)
		} else {
			return 1<<39 | (oak2&mask)<<13 | (kicker & mask)
		}
	} else {
		return kicker
	}
}
