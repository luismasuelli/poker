package seats

import (
	"github.com/luismasuelli/poker-go/engine/games/messages/tables"
	"github.com/luismasuelli/poker-go/engine/games/tables/seats"
)

// Messages related to a seat will hold their
// game ID, table ID, and Seat ID. This is only
// meaningful when a user is actually observing
// a table status, and the seat belongs to that
// table.
type SeatMessage struct {
	tables.TableMessage
	Seat uint8
}

// Tells when a seat has been occupied by
// a new player (it only shows its display
// data), and tells the new seat's stack.
type SeatHasBeenOccupied struct {
	SeatMessage
	PlayerDisplay interface{}
	Stack         uint64
}

// Tells when a seat has been released.
// It also tells reason and arbitrary
// description data.
type SeatHasBeenReleased struct {
	SeatMessage
}

// Tells when a seat's status was changed.
type SeatStatusHasChanged struct {
	SeatMessage
	PlayerDisplay interface{}
	NewStatus     seats.Status
}

// Tells when a seat's flag was set.
type SeatFlagHasBeenSet struct {
	SeatMessage
	Flag       seats.Flags
	FinalFlags seats.Flags
}

// Tells when a seat's flag was cleared.
type SeatFlagHasBeenCleared struct {
	SeatMessage
}

// Tells when a seat stack gained chips.
type SeatStackHasGrown struct {
	SeatMessage
	Chips      uint64
	FinalStack uint64
}

// Tells when a seat stack lost chips.
type SeatStackHasShrank struct {
	SeatMessage
	Chips      uint64
	FinalStack uint64
}

// Tells when a seat stack was set.
// This will seldom be used, save for
// tournament initialization or perhaps
// administrative actions.
type SeatStackHasChanged struct {
	SeatMessage
	FinalStack uint64
}
