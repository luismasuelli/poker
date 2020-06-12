package chips

// Chip stacks serve to tell many different
// things in a table or even lobby:
// - The player's total chips.
// - The player's transient pot chips.
// - The table pot(s).
// For this to work, stacks will have different
// methods:
// - Give a neutral representation of the chips.
// - Give a formatted representation of the chips.
// - Create a stack of the same type, with a
//   given amount of chips.
//
// On a given poker game (regardless the number
// of involved tables), only ONE type of stack
// will be used.
type ChipStack interface {
	Amount() uint64
	FormattedAmount() string
	New(amount uint64) ChipStack
}
