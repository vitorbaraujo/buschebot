package repliers

import (
	"fmt"
	"os"

	"github.com/vitorbaraujo/buschebot/reply"
)

type BuscheReplier struct{}

func init() {
	fmt.Println("Registering busche replier")
	reply.RegisterReplier(&BuscheReplier{})
}

func (BuscheReplier) Reply(msg string) string {
	if !reply.IsQuestion(msg) {
		return msg + "?"
	}

	return ""
}

func (BuscheReplier) CanReadMessage(payload *reply.MessagePayload) bool {
	buscheId, ok := os.LookupEnv("BUSCHE_ID")
	return ok && payload.UserId == buscheId
}