package players

import "github.com/luismasuelli/poker-go/engine"

// This is a contract to notify players anything that
// happens in the poker side of the game.
type Notifiable interface {
	// Notifies the player about any event that could
	// occur in a game, a particular table or lobby,
	// and/or perhaps a particular seat.
	//
	// Also, notifications can be sent regarding a
	// particular request from / on behalf of the
	// player.
	//
	// The data to send involves a code and optional
	// keyword arguments, and the notify operation
	// must not block for long time.
	Notify(gameID engine.GameID, tableID engine.TableID, seatID engine.SeatID, requestID engine.RequestID,
		   code string, data map[string]interface{})
}
