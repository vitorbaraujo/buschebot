package reply

import (
	"fmt"
	"strings"

	"github.com/vitorbaraujo/buschebot/storage"
)

func IsQuestion(msg string) bool {
	return strings.HasSuffix(msg, "?")
}

func IsIndagation(msg string) bool {
	if !IsQuestion(msg) {
		return false
	}

	data, err := storage.GetData()
	if err != nil {
		fmt.Errorf("getting data from storage")
		return false
	}

	for _, prefix := range data.IndagationPrefixes {
		if strings.HasPrefix(msg, prefix) {
			return true
		}
	}

	return false
}
