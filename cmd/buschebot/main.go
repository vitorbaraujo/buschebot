package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vitorbaraujo/buschebot/bot"
	"github.com/vitorbaraujo/buschebot/repliers"
	"github.com/vitorbaraujo/buschebot/reply"
)

func main() {
	var apiToken string
	if val, ok := os.LookupEnv("BOT_API_TOKEN"); ok {
		apiToken = val
	}

	tgbot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Panicf("creating tgbot: %v", err)
	}

	tgbot.Debug = true
	log.Printf("Authorized on account %v", tgbot.Self.UserName)

	cli := bot.NewClient(tgbot, []reply.Replier{
		&repliers.BuscheReplier{},
	})

	if err := cli.Run(); err != nil {
		log.Panicf("querying messages: %v", err)
	}
}
