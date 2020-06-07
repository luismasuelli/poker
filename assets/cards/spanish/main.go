package french

var faces = [64]string{
	"1c", "2c", "3c", "4c", "5c", "6c", "7c", "8c", "9c", "Sc", "Cc", "Rc",
	"1o", "2o", "3o", "4o", "5o", "6o", "7o", "8o", "9o", "So", "Co", "Ro",
	"1b", "2b", "3b", "4b", "5b", "6b", "7b", "8b", "9b", "Sb", "Cb", "Rb",
	"1e", "2e", "3e", "4e", "5e", "6e", "7e", "8e", "9e", "Se", "Ce", "Re",
	"*w", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!",
	"!!", "!!", "!!", "!!",
}

// Defines a spanish card (1 out of 49, since the
// wildcards also count).
type Card uint8

func (card Card) Face() string {
	return faces[card]
}

func (card Card) Set() uint64 {
	return (1 << 49) - 1
}
