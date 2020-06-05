package rand

import (
	"math/rand"
	"github.com/luismasuelli/poker/assets/cards"
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

// A custom shuffler option for the NewShuffler
// constructor.
type ShufflerOption func(shuffler *CustomShuffler)

// Option to create a shuffler that always seeds
// on each shuffle, using the current time.
func TimeSeedOnShuffle(shuffler *CustomShuffler) {
	shuffler.timeSeed = true
}

// Option to set the initial seed of the shuffler.
func InitSeedWith(seed int64) ShufflerOption {
	return func(shuffler *CustomShuffler) {
		shuffler.randObj.Seed(seed)
	}
}

// Creates a new rand-based shuffler, with option to
// choose a rand-standard source (using nil as source)
// or a custom one. Additional options may be given
// to set the initial source's seed and/or tell the
// shuffler to continuously seed on current time.
func NewShuffler(source rand.Source, options ...ShufflerOption) *CustomShuffler {
	if source == nil {
		source = rand.NewSource(1)
	}
	shuffler := &CustomShuffler{
		randObj: rand.New(source),
		timeSeed: false,
	}
	for _, option := range options {
		option(shuffler)
	}
	return shuffler
}