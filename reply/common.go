package reply

import "strings"

func IsQuestion(msg string) bool {
	return strings.HasSuffix(msg, "?")
}

func IsIndagation(msg string) bool {
	if !IsQuestion(msg) {
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
