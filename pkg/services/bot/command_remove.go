package bot

import (
	"context"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"service/pkg/entities"
)

const removeCommandValue = "remove"

func (b *Bot) removeCommand(ctx context.Context, args string, chatID int64) (err error) {

	requester := entities.NewRequester(strconv.FormatInt(chatID, 10))

	err = b.actions.DeleteLink(ctx, args, requester)
	if err != nil {
		return
	}

	reply := tgbotapi.NewMessage(chatID, b.buildRemoveResponse())
	_, err = b.botAPI.Send(reply)
	return
}

func (b *Bot) buildRemoveResponse() string {
	return linkRemovedMessage
}
