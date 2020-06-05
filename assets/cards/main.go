package cards

// Cards are nothing by themselves, save for their
// visual representation. This is done in their
// "face" method. The "face" should be matched in
// front-end with an appropriate representation.
type Card interface {
	Face() string
}
