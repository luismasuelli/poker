// Actors here are only defined as interfaces. Since this engine
// is meant to be used in an unrestricted diversity of environments,
// the only requirement is that the end-users of this architecture
// be interfaces to be implemented. Then, the concrete classes that
// implement these interfaces should not matter at all to this
// architecture, which is useful to mock/test and to implement them
// in said diversity of environments (e.g. a stand-alone poker game
// or part of another, bigger, game or application).
package actors
