package messages

type Message struct {
	// Message content.
	Body       interface{}

	// Some messages can be in response to other
	// messages sent by client. A value of 0
	// means this message is not replying.
	ReplyingTo uint64

	// Some messages can have an underlying reason
	// to do what they do (e.g. administrative
	// actions).

	// Reason code for the message, if used.
	Reason     string

	// Reason details for the message, if used.
	Details    interface{}
}
