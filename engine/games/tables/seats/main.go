package seats

import "github.com/luismasuelli/poker-go/engine/players"

// Status of a seat during a game.
type Status uint32

const (
	// No player is occupying this seat.
	Free Status = iota
	// The player is waiting the next hand
	// in which they "will chose to play".
	Waiting
	// The player is playing the hand, and
	// has $X > 0.
	Active
	// The player is playing the hand, but
	// it has $0, which marks them as being
	// "all-in".
	AllIn
	// The player folded.
	Folded
)

// Additional flags to use (e.g. "sit out")
// to complement a status.
type Flags uint32

const (
	// No flags at all.
	Nothing Flags = 0
	// The seat is marked as "sit out".
	SitOut Flags = 1
)

// Seats are managed by the tables. They
// keep a reference to the player and to
// each aspect of their participation in
// the table.
type Seat struct {
	player players.Player
	status Status
	flags  Flags
	stack  uint64
}

// Gets the underlying sit player.
func (seat *Seat) Player() players.Player {
	return seat.player
}

// Gets the stack of this seat.
func (seat *Seat) Stack() uint64 {
	return seat.stack
}

// Returns the status of a seat. There are
// only 5 statuses in poker games, although
// flags may condition when and how these
// states are changed.
func (seat *Seat) Status() Status {
	return seat.status
}

// Returns the flags of this seat. So far,
// this only involves the "sit out" button.
func (seat *Seat) Flags() Flags {
	return seat.flags
}

// Tries to sit a player, if given and with
// a non-empty stack, and only if no player
// is already sit. It is recommended for
// this method to be redefined on composition
// if more stuff is to be changed on sit.
func (seat *Seat) Sit(player players.Player, stack uint64) bool {
	if seat.status == Free || player == nil || stack == 0 {
		return false
	} else {
		seat.player = player
		seat.stack = stack
		seat.status = Waiting
		return true
	}
}

// Pops a player and stack from the seat. It
// returns (nil, 0) if the seat was empty.
// It is recommended for this method to be
// redefined on composition if more stuff is
// to be changed on pop.
func (seat *Seat) Pop() (player players.Player, stack uint64) {
	player = seat.player
	stack = seat.stack
	seat.player = nil
	seat.stack = 0
	seat.status = Free
	seat.flags = Nothing
	return
}

// Sets the status of this seat. The "free"
// status cannot set, and cannot replaced,
// by this method.
func (seat *Seat) SetStatus(status Status) bool {
	if seat.player == nil || status == Free {
		// Cannot set the status of a table
		// without player, and cannot set
		// the status of a table with player
		// to Free.
		return false
	} else {
		seat.status = status
		return true
	}
}

// Sets a flag on this seat. This can only
// be done to non-free seats.
func (seat *Seat) SetFlag(flag Flags) bool {
	if seat.player == nil {
		return false
	} else {
		seat.flags |= flag
		return true
	}
}

// Clears a flag on this seat. This can only
// be done to non-free seats.
func (seat *Seat) ClearFlag(flag Flags) bool {
	if seat.player == nil {
		return false
	} else {
		seat.flags &= ^flag
		return true
	}
}

// Takes chips from the seat. This can only
// be done to non-free seats that can afford
// the specified chips amount.
func (seat *Seat) TakeChips(chips uint64) bool {
	if seat.player == nil || seat.stack < chips {
		return false
	} else {
		seat.stack -= chips
		return true
	}
}

// Gives chips to the seat. This can only be
// done to non-free seats that can receive
// the amount without overflowing. It is
// recommended that table levels and user
// accounts are defensively restricted to
// reach a case like this.
func (seat *Seat) GiveChips(chips uint64) bool {
	if seat.player == nil || seat.stack > (^(uint64(0) - chips)) {
		return false
	} else {
		seat.stack += chips
		return true
	}
}

// Sets the chips to the seat. This can only
// be done to non-free seats, and care should
// be taken when calling this method, since it
// may cause overflow as well when receiving
// more money or popping the seat in a cash
// 1-table game.
func (seat *Seat) SetStack(chips uint64) bool {
	if seat.player == nil {
		return false
	} else {
		seat.stack = chips
		return true
	}
}