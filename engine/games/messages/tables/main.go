package tables

import "github.com/luismasuelli/poker-go/engine/games/messages/games"

// Messages related to a table will hold their
// game ID and table ID. This is only meaningful
// when a user is actually observing a table
// status.
type TableMessage struct {
	games.GameMessage
	TableID uint32
}
