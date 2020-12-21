package repliers

import (
	"os"

	"github.com/vitorbaraujo/buschebot/reply"
)

// BuscheReplier is a custom replier that deals with messages coming from the telegram user "jpbusche".
type BuscheReplier struct{}

// Reply answers a message from jpbusche.
func (BuscheReplier) Reply(msg string) string {
	if !reply.IsQuestion(msg) {
		return msg + "?"
	}

	return ""
}

// CanReadMessage determines whether this replier can answer a specific message.
// This replier can read a message only if it's coming from jpbusche and it's telegram ID is configured.
func (BuscheReplier) CanReadMessage(payload *reply.MessagePayload) bool {
	buscheID, ok := os.LookupEnv("BUSCHE_ID")
	return ok && payload.UserID == buscheID
}
