package engine

// There are many different games that start and end, and will never
// be created again (but perhaps different instances of them) like
// tournaments, and there are other games that never end, but can
// be shutdown and recreated (with the same identity) like cash tables.
// This said, there is a notion of key or identity for games.
type GameID interface{}

// All the games have at least one table. A typical game will have
// few tables, at least 1. Multi-table tournaments can have thousands
// of tables instead. TableID 0xffffffff is reserved for lobby in MTT.
// TableID 0 means "no table in particular".
type TableID uint32

// Tables will have a small number of seats. SeatID 0 means "no seat
// in particular".
type SeatID uint8

// Player requests will be big numbers to enumerate transactions that
// can be tracked in future messages. RequestID 0 means "no request
// in particular" and should never be used, to avoid confusion.
type RequestID uint64

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

// Assets are registered under a special registry (specially for
// tickets, which will be a lot, with custom caption and for
// several different tournaments and promos).
type AssetID string