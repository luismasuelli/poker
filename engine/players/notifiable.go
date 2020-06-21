package players

// This is a contract to notify players anything that
// happens in the poker side of the game.
type Notifiable interface {
	// Notifies the player about any event that could
	// occur in a game, a particular table or lobby,
	// and/or perhaps a particular seat.
	//
	// The data to send involves a code and optional
	// keyword arguments, and the notify operation
	// must not block for long time.
	Notify(message interface{})
}
