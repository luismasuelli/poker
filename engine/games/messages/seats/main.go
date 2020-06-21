package seats

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
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

// Tells when a seat received cards,
// which may be revealed or nil. This
// message stands both for initial
// cards and additional drawn cards
// for "draw" modes.
type SeatDrewCards struct {
	SeatMessage
	Cards []cards.Card
}

// Tells when a seat gives N cards.
// Intended for "draw" modes only.
type SeatGaveCards struct {
	SeatMessage
	Count int
}

// Tells when client's seat received
// cards, which will always be revealed
// but with the "hidden/revealed" flags
// to know how to hint the cards as
// being available to the client but
// hidden to others.
type YouDrewCards struct {
	SeatMessage
	Cards []cards.Card
	Shown []bool
}

// Tells when client's seat gave some
// cards, and which indices in particular.
// Intended for "draw" games.
type YouGaveCards struct {
	Indices []int
}
