// A simple anonymizer bot that echoes back any message it gets in a private chat, so that you can forward it without exposing the original sender.
package main

import (
	"context"
	"flag"
	"sync"

	"github.com/sdurz/axon"
	"github.com/sdurz/ubot"
)

var (
	apiKey string
)

func init() {
	flag.StringVar(&apiKey, "apikey", "", "api key")
	flag.Parse()
}

func main() {
	var (
		wg  sync.WaitGroup
		bot *ubot.Bot
	)
	bot = ubot.NewBot(&ubot.Configuration{APIToken: apiKey})
	bot.AddMessageHandler(ubot.MessageIsPrivate, func(ctx context.Context, b *ubot.Bot, message axon.O) (done bool, err error) {
		messageId, _ := message.GetInteger("message_id")
		chatId, _ := message.GetInteger("chat.id")
		bot.CopyMessage(axon.O{
			"chat_id":      chatId,
			"from_chat_id": chatId,
			"message_id":   messageId,
		})
		return
	})
	bot.Forever(context.Background(), &wg, ubot.GetUpdatesSource)
}
