package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

// Data represents the data used by the bot to answer messages.
type Data struct {
	IndagationPrefixes []string `json:"indagation_prefixes"`
	Responses          []string `json:"responses"`
	ShortAnswers       []string `json:"short_answers"`
}

var data *Data

// DataFile is the name of the file that contains the data used by the bot.
var DataFile = "storage/data.json"

// GetData retrieves the data from the database.
// Currently, it fetches data from a local json file.
func GetData() (*Data, error) {
	if data != nil {
		return data, nil
	}

	file, err := os.Open(DataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	return data, err
}

// GetRandomResponse gets a random response from the database.
func GetRandomResponse() (string, error) {
	data, err := GetData()
	if err != nil {
		return "", fmt.Errorf("getting random response: %w", err)
	}

	// nolint:gosec
	index := rand.Int() % len(data.Responses)
	return data.Responses[index], nil
}

// GetRandomShortAnswer gets a random short answer from the database.
func GetRandomShortAnswer() (string, error) {
	data, err := GetData()
	if err != nil {
		return "", fmt.Errorf("getting random short answer: %w", err)
	}

	// nolint:gosec
	index := rand.Int() % len(data.ShortAnswers)
	return data.ShortAnswers[index], nil
}
