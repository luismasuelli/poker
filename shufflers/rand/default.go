package rand

import (
	"github.com/luismasuelli/poker-go/rules"
	"math/rand"
	"time"
)

// A default shuffler uses the default, global, rand
// functions for Seed and Shuffle. This works similar
// to the custom shuffler, except that it works with
// the global source, which is safe in concurrency.
type DefaultShuffler struct {
	timeSeed bool
}

// Shuffles a deck using its Len and Swap methods
// in the underlying rand object.
func (shuffler *DefaultShuffler) Shuffle(deck rules.Deck) {
	if shuffler.timeSeed {
		rand.Seed(time.Now().UTC().UnixNano())
	}
	rand.Shuffle(deck.Len(), deck.Swap)
}

// Creates a new rand-based Shuffler, which uses the
// standard Seed and Shuffle functions, which may seed
// with the current time on each shuffle, or not.
func NewDefaultShuffler(timeSeed bool) *DefaultShuffler {
	return &DefaultShuffler{timeSeed}
}
