package messages

type Message struct {
	// Some messages can have an underlying reason
	// to do what they do (e.g. administrative
	// actions).
	Reason  string
	Details interface{}
}
