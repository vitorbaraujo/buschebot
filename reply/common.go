package reply

import (
	"strings"

	"github.com/vitorbaraujo/buschebot/storage"
)

// IsQuestion checks if a message is a question.
// A question is a sentence that ends with a question mark.
func IsQuestion(msg string) bool {
	return strings.HasSuffix(msg, "?")
}

// IsIndagation checks if a message is an indagation.
// An indagation is a question that contains some predefined keyword prefixes.
// See storage/data.json.
func IsIndagation(msg string) bool {
	if !IsQuestion(msg) {
		return false
	}

	data, err := storage.GetData()
	if err != nil {
		// error getting data from storage
		return false
	}

	for _, prefix := range data.IndagationPrefixes {
		if strings.HasPrefix(msg, prefix) {
			return true
		}
	}

	return false
}
