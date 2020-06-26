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

// A buy-in differs from a payment that it may
// allow a range of amounts.
type BuyIn struct {
	asset     Asset
	minAmount uint64
	maxAmount uint64
}

// Gets the asset of this buy-in.
func (buyIn *BuyIn) Asset() Asset {
	return buyIn.asset
}

// Gets the minimum amount of this buy-in.
func (buyIn *BuyIn) MinAmount() uint64 {
	return buyIn.minAmount
}

// Gets the maximum amount of this buy-in.
func (buyIn *BuyIn) MaxAmount() uint64 {
	return buyIn.maxAmount
}

// Creates a new buy-in by giving the asset and
// the allowed range of amounts (typically, only
// cash games will allow ranges - tournaments
// will only allow a single value).
func NewBuyIn(asset Asset, minAmount, maxAmount uint64) *BuyIn {
	if minAmount > maxAmount {
		minAmount, maxAmount = maxAmount, minAmount
	}
	return &BuyIn{asset, minAmount, maxAmount}
}

// A map of assets against the buy-ins.
type BuyIns map[Asset]BuyIn