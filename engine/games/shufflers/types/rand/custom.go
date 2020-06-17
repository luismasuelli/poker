package rand

import (
	"github.com/luismasuelli/poker-go/engine/games/cards"
	"math/rand"
	"time"
)

// A custom shuffler takes a random source to
// perform the shuffle. The source will come
// with its own seed state.
type CustomShuffler struct {
	randObj  *rand.Rand
	timeSeed bool
}

// Shuffles a deck using its Len and Swap methods
// in the underlying rand object.
func (shuffler *CustomShuffler) Shuffle(deck cards.Deck) {
	if shuffler.timeSeed {
		shuffler.randObj.Seed(time.Now().UTC().UnixNano())
	}
	shuffler.randObj.Shuffle(deck.Len(), deck.Swap)
}

// Creates a new rand-based shuffler, with option to
// choose a rand-standard source (using nil as source)
// or a custom one. Additional options will be given
// to set the initial source's seed and/or tell the
// shuffler to continuously seed on current time.
func NewShuffler(source rand.Source, timeSeed bool, initialSeed int64) *CustomShuffler {
	if source == nil {
		source = rand.NewSource(initialSeed)
	}
	shuffler := &CustomShuffler{
		randObj:  rand.New(source),
		timeSeed: timeSeed,
	}
	return shuffler
}
