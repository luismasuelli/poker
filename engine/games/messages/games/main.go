package games

// Messages related to a game will hold their
// ID. This is only meaningful when a user is
// actually observing a game  status.
type GameMessage struct {
	// The ID of the game this content is
	// related to.
	GameID  interface{}

	// Content of the message.
	Content interface{}
}
