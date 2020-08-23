package reply

type DefaultReplier struct{}

func (DefaultReplier) Reply(string) string {
	return ""
}

func (DefaultReplier) CanReadMessage(*MessagePayload) bool {
	return true
}



