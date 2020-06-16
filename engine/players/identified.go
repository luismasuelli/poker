package players

// Players have an identification
// that can be used to refer them
// in several commands. Identifiers
// should be unique across all the
// registered players in the engine.
type Identified interface {
	Identification() interface{}
}
