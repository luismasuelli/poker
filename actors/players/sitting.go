package players

// Sitting players correspond to the contract a
// player satisfies while in a seat. Sitting
// players only exist in the context of a game
// in particular: game engines must ensure these
// sitting players are not shared (among games)
// and also that they occupy a SINGLE seat inside
// them (this means: they are not duplicated in
// a table or across tables in the same game).
type SittingPlayer interface {
	// A reference to the player. This is in order
	// to back-refer it, e.g. to award or refund
	// it when leaving the table. Another reason
	// to back-refer the player is to get the set
	// of display data (which is arbitrary and
	// implementation-specific).
	Player() Player
	// The current chips stack in this tournament.
	// The initial state of a sitting player,
	// before buy-in or chips assignment after
	// registration, should be Chips() == 0.
	Chips() uint64
	// Adds chips (e.g. on buy-in, assignment, or
	// winning) to the player. THIS METHOD SHOULD
	// PANIC ON OVERFLOW, and buy-in requests and
	// tournament assignments should account for
	// the limit of uint64 in a defensively coded
	// fashion.
	AddChips(uint64) uint64
	// Takes chips from the player. Returns true
	// if it could take the chips, and returns
	// false if this sitting player cannot afford
	// that (local, in-game) amount.
	TakeChips(uint64) bool
}
