package tables

import "github.com/luismasuelli/poker-go/engine/games/rules/showdowns"

// Table messages will be related to a
// table (in a particular game - this
// table message will be wrapped in a
// game message).
type TableMessage struct {
	// The table ID.
	TableID uint32

	// The content of this table message.
	Content interface{}
}

// Tells when a showdown will occur or
// be skipped.
type Showdown struct {
	HandID  uint64
	Mode    showdowns.Mode
	Skipped bool
}