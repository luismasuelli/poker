package games

// Messages related to a game will hold their
// ID. This is only meaningful when a user is
// actually observing a game  status.
type GameMessage struct {
	GameID interface{}
}
