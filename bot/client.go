package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vitorbaraujo/buschebot/reply"
)

// Client is a bot client, which holds a reference to the Telegram bot API and a list of custom repliers.
type Client struct {
	bot         *tgbotapi.BotAPI
	replyClient *reply.Client
}

// NewClient creates a new bot client.
func NewClient(bot *tgbotapi.BotAPI, customRepliers []reply.Replier) *Client {
	return &Client{
		bot:         bot,
		replyClient: reply.NewClient(customRepliers),
	}
}

// Run queries telegram messages and answers them, if possible.
func (c *Client) Run() error {
	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 60

	updates, err := c.bot.GetUpdatesChan(cfg)
	if err != nil {
		return err
	}

	// This runs indefinitely.
	for update := range updates {
		if update.Message == nil {
			continue
		}

		response, err := c.replyClient.GetReply(&reply.MessagePayload{
			Text:   update.Message.Text,
			UserID: fmt.Sprint(update.Message.From.ID),
		})
		if err != nil {
			return err
		}

		if response.Text == "" {
			// bot did not come up with an answer.
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.Text)
		if response.Reply {
			msg.ReplyToMessageID = update.Message.MessageID
		}

		if _, err := c.bot.Send(msg); err != nil {
			// failed to send message, continue to next response
			continue
		}
	}

	return nil
}
