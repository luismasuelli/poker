package cards

// Cards, from a given set, will only have their
// "face" (a precomputed string) and a method to
// give a "hint" of the existing cards in the set,
// in an uint64 result (Cards will belong to a set
// of at most 64 cards).
type Card interface {
	Face() string
	Set() uint64
}
