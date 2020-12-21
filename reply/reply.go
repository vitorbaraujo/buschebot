package reply

import (
	"math/rand"
	"strings"

	"github.com/vitorbaraujo/buschebot/storage"
)

// Replier is an interface responsible for replying to certain messages.
type Replier interface {
	Reply(string) string
	CanReadMessage(payload *MessagePayload) bool
}

// Client is a reply client.
type Client struct {
	repliers []Replier
}

// Response defines a bot message response.
type Response struct {
	// Text is the reply message.
	Text string
	// Reply tells whether or not reply to the message.
	// See https://telegram.org/blog/replies-mentions.
	Reply bool
}

// MessagePayload defines a telegram message.
type MessagePayload struct {
	// Text is the content of the message.
	Text string
	// UserID contains the id of the message sender.
	UserID string
}

// RandInt is a function that returns a random int. It's exposed to allow customization in tests.
var RandInt = rand.Int

// NewClient creates a new reply client.
func NewClient(customRepliers []Replier) *Client {
	return &Client{
		repliers: customRepliers,
	}
}

// GetReply parses a message and responds accordingly.
// It can differ between a regular question and an indagation (see isIndagation method).
func (c *Client) GetReply(payload *MessagePayload) (*Response, error) {
	text := strings.ToLower(strings.TrimSpace(payload.Text))
	answer := c.getReplier(payload).Reply(text)
	var err error

	if answer == "" && IsQuestion(text) {
		if IsIndagation(text) {
			answer, err = storage.GetRandomResponse()
			if err != nil {
				return nil, err
			}
		} else {
			answer, err = storage.GetRandomShortAnswer()
			if err != nil {
				return nil, err
			}
		}
	}

	return &Response{
		Text:  answer,
		Reply: RandInt()%2 == 0,
	}, nil
}

func (c *Client) getReplier(msg *MessagePayload) Replier {
	for _, replier := range c.repliers {
		if replier.CanReadMessage(msg) {
			return replier
		}
	}

	return &DefaultReplier{}
}
