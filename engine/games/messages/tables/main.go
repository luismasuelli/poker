package tables

// Table messages will be related to a
// table (in a particular game - this
// table message will be wrapped in a
// game message).
type TableMessage struct {
	// The table ID.
	TableID uint32

	// The content of this table message.
	Content interface{}
}
