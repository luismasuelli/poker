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

// A payment is an amount of a specific asset.
// Payments may come in two different cases,
// being prizes and buy-ins.
type Payment struct {
	asset  Asset
	amount uint64
}

// Gets the payment asset.
func (payment *Payment) Asset() Asset {
	return payment.asset
}

// Gets the payment amount.
func (payment *Payment) Amount() uint64 {
	return payment.amount
}

// Creates a new payment by giving the asset and
// the amount. Nil values will be dealt with later.
func NewPayment(asset Asset, amount uint64) *Payment {
	return &Payment{asset, amount}
}

// Prizes are intended for tournaments and specified
// in order. This means: [0] will contain the prize
// for the 1st place, and [9] for the 10th place.
type Prizes []Payment

// Buy-ins are a mapping of a required asset and its
// required amount.
type BuyIns map[Asset]uint8