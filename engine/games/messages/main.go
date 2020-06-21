package messages

type Message struct {
	// Some messages can have an underlying reason
	// to do what they do (e.g. administrative
	// actions).
	Reason  string
	Details interface{}
	// Some messages can be in response to other
	// messages sent by client. A value of 0
	// means this message is not replying-
	ReplyingTo uint64
}
