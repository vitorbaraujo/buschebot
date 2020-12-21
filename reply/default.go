package reply

// DefaultReplier is a replier that can read and answer any messages.
type DefaultReplier struct{}

// Reply answer a message with an empty string.
func (DefaultReplier) Reply(string) string {
	return ""
}

// CanReadMessage determiner whether or not this replier can read a message.
// The default replier can always read any message.
func (DefaultReplier) CanReadMessage(*MessagePayload) bool {
	return true
}
