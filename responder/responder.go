package responder

import (
	"math/rand"
	"strings"
)

// Response defines a bot message response.
type Response struct {
	// Text is the reply message.
	Text string
	// Reply tells whether or not reply to the message.
	// See https://telegram.org/blog/replies-mentions.
	Reply bool
}

// RandInt is a function to return a random int. It's exposed to allow customization in tests.
var RandInt = rand.Int

// ReplyMessage parses a message and responds accordingly.
// It can differ between a regular question and an indagation (see isIndagation method).
func ReplyMessage(msg string) *Response {
	var answer string
	msg = strings.ToLower(strings.TrimSpace(msg))

	if !isQuestion(msg) {
		return &Response{}
	}

	if isIndagation(msg) {
		answer = "sei la"
	} else {
		answer = "sim"
	}

	resp := &Response{Text: answer}
	resp.Reply = RandInt() % 2 == 0
	return resp
}

func isQuestion(msg string) bool {
	return strings.HasSuffix(msg, "?")
}

func isIndagation(msg string) bool {
	if !isQuestion(msg) {
		return false
	}

	prefixes := []string{
		"qual",
		"o que",
		"quando",
		"quem",
		"pq",
		"porque",
		"por que",
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(msg, prefix) {
			return true
		}
	}

	return false
}