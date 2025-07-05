package client

// SMSProvider defines the interface for sending SMS across multiple platforms
// SendMessage returns a slice of SendResponse, one per recipient.
type SMSProvider interface {
	SendMessage(ref string, recipients []string, content, lang string) ([]SendResponse, error)
}

// SendResponse is a common response for SendMessage
type SendResponse struct {
	Recipient string // destination phone number
	JobID     string // platform-specific message ID
}
