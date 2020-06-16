package assets

// There are two asset types: currencies, and tickets.
// Both can be used as buy-ins or prizes in tournaments,
// but only currencies can be used for cash tables,
// since the buy-in maps 1:1 to in-game chips, and 1:1
// with the refund of the table.
type AssetType uint8

const (
	Currency AssetType = iota
	Ticket
)

// Assets have a type, a caption, and an ID. Captions
// are specially useful when dealing with tickets.
type Asset interface {
	ID()      interface{}
	Type()    AssetType
	Caption() string
}
