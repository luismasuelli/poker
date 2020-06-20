package games

import "github.com/luismasuelli/poker-go/engine/games/messages"

// Messages related to a game will hold their
// ID. This is only meaningful when a user is
// actually observing a game  status.
type GameMessage struct {
	messages.Message
	GameID interface{}
}
